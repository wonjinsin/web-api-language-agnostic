package model

import (
	"pikachu/util"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// User is model for user
type User struct {
	UID       string                `json:"uid" gorm:"primaryKey"`
	Email     string                `json:"email"`
	Password  Password              `json:"password,omitempty"`
	Name      *string               `json:"name,omitempty"`
	CreatedAt time.Time             `json:"createdAt" gorm:"<-:create;autoCreateTime;not null"`
	UpdatedAt time.Time             `json:"updatedAt" gorm:"autoUpdateTime;not null"`
	DeletedAt soft_delete.DeletedAt `json:"deletedAt" gorm:"default:0"`
}

// NewUserBySignup is for new user by signup
func NewUserBySignup(su *Signup) (user *User, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(su.Password), 14)
	if err != nil {
		return nil, err
	}

	return &User{
		UID:      uuid.New().String(),
		Email:    su.Email,
		Password: Password(bytes),
		Name:     su.Name,
	}, nil
}

// AfterFind is for after find
func (u *User) AfterFind(tx *gorm.DB) (err error) {
	ctx := tx.Statement.Context
	login, _ := ctx.Value(util.LoginKey).(bool)
	if !login {
		u.Password = ""
	}
	return
}
