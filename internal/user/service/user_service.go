package service

import (
	"context"
	"time"

	"user-profile-service/internal/config"
	"user-profile-service/internal/domain"
	"user-profile-service/internal/rabbitmq"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo      domain.UserRepository
	publisher *rabbitmq.Publisher
	cfg       *config.Config
}

func NewUserService(repo domain.UserRepository, publisher *rabbitmq.Publisher, cfg *config.Config) domain.UserService {
	return &userService{repo: repo, publisher: publisher, cfg: cfg}
}

func (s *userService) Register(ctx context.Context, req domain.RegisterRequest) (*domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:           uuid.New(),
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Phone:        req.Phone,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Role:         req.Role,
		Status:       "active",
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Login(ctx context.Context, req domain.LoginRequest) (*domain.LoginResponse, error) {
	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	if user.Status == "suspended" || user.Status == "banned" {
		return nil, domain.ErrUserSuspendedBanned
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	// Issue standard stateless JWT token signed by symmetrical key secret
	claims := domain.JWTClaims{
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "iam-user-service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(s.cfg.JWTSecret))
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{Token: tokenStr, User: user}, nil
}

func (s *userService) GetProfile(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *userService) UpdateProfile(ctx context.Context, id uuid.UUID, req domain.UpdateUserRequest) (*domain.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}

	if err := s.repo.Update(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *userService) ChangeUserStatus(ctx context.Context, id uuid.UUID, req domain.BanUserRequest) (*domain.User, error) {
	user, err := s.repo.UpdateStatus(ctx, id, req.Status)
	if err != nil {
		return nil, err
	}

	// Intercept ban transition phase to emit downstream notification messages
	if user.Status == "banned" {
		event := domain.UserEventPayload{
			UserID:    user.ID.String(),
			Email:     user.Email,
			Role:      user.Role,
			Status:    user.Status,
			Timestamp: time.Now().UTC(),
		}
		_ = s.publisher.PublishEvent(ctx, "user.banned", event)
	}

	return user, nil
}
