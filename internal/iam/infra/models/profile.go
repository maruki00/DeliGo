package models

import (
	"time"
)

type Profile struct {
	ID        string
	UserID    string
	FullName  string
	Avatar    string
	Bio       string
	DeletedAt *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (p *Profile) SetID(ID string) {
	p.ID = ID
}
func (p *Profile) SetUserID(UserID string) {
	p.UserID = UserID
}
func (p *Profile) SetFullName(FullName string) {
	p.FullName = FullName
}
func (p *Profile) SetAvatar(Avatar string) {
	p.Avatar = Avatar
}
func (p *Profile) SetBio(Bio string) {
	p.Bio = Bio
}
func (p *Profile) SetDeletedAt(DeletedAt *time.Time) {
	p.DeletedAt = DeletedAt
}
func (p *Profile) SetCreatedAt(CreatedAt *time.Time) {
	p.CreatedAt = CreatedAt
}
func (p *Profile) SetUpdatedAt(UpdatedAt *time.Time) {
	p.UpdatedAt = UpdatedAt
}

func (p *Profile) GetID() string {
	return p.ID
}
func (p *Profile) GetUserID() string {
	return p.UserID
}
func (p *Profile) GetFullName() string {
	return p.FullName
}
func (p *Profile) GetAvatar() string {
	return p.Avatar
}
func (p *Profile) GetBio() string {
	return p.Bio
}
func (p *Profile) GetDeletedAt() *time.Time {
	return p.DeletedAt
}
func (p *Profile) GetCreatedAt() *time.Time {
	return p.CreatedAt
}
func (p *Profile) GetUpdatedAt() *time.Time {
	return p.UpdatedAt
}
