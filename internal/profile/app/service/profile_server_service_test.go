package service

import (
	"context"
	"testing"

	"github.com/maruki00/deligo/internal/profile/domain/contract"
	"github.com/maruki00/deligo/internal/profile/domain/entity"
	profile_grpc "github.com/maruki00/deligo/internal/profile/infra/grpc/profile"
	"github.com/maruki00/deligo/internal/profile/infra/model"
)

type fakeProfileRepo struct {
	updatedID     string
	updatedFields map[string]any
}

func (f *fakeProfileRepo) Save(context.Context, entity.ProfileEntity) error { return nil }
func (f *fakeProfileRepo) Disable(context.Context, string) error            { return nil }
func (f *fakeProfileRepo) FindByUserID(context.Context, string) (*model.Profile, error) {
	return nil, nil
}
func (f *fakeProfileRepo) FindByID(context.Context, string) (*model.Profile, error) { return nil, nil }
func (f *fakeProfileRepo) Update(_ context.Context, id string, fields map[string]any) error {
	f.updatedID = id
	f.updatedFields = fields
	return nil
}
func (f *fakeProfileRepo) UpdateAvatar(context.Context, string, string) error { return nil }

var _ contract.IPorofileRepository = (*fakeProfileRepo)(nil)

func TestUpdateForwardsFieldsToRepository(t *testing.T) {
	repo := &fakeProfileRepo{}
	svc := NewProfileService(repo)

	resp, err := svc.Update(context.Background(), &profile_grpc.UpdateProfileRequest{
		ID: "1234",
		Fields: map[string]string{
			"full_name": "Jane Doe",
		},
	})
	if err != nil {
		t.Fatalf("update returned error: %v", err)
	}
	if resp == nil {
		t.Fatal("expected response to be returned")
	}
	if repo.updatedID != "1234" {
		t.Fatalf("expected repository id to be updated, got %q", repo.updatedID)
	}
	if repo.updatedFields["full_name"] != "Jane Doe" {
		t.Fatalf("expected full_name field to be forwarded, got %v", repo.updatedFields["full_name"])
	}
}
