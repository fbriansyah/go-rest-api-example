package domain

import (
	"api-example/pkg/util"
	"time"

	"github.com/google/uuid"
)

// UserStatus membuat type alias string
type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"fullname"`
	Email    string    `json:"email"`
	// Password  hashed version
	Password  string     `json:"password"`
	Status    UserStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func NewUser(fullname, email string) *User {
	return &User{
		ID:        uuid.New(),
		FullName:  fullname,
		Email:     email,
		Status:    UserStatusActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u *User) IsActive() bool {
	return u.Status == UserStatusActive
}

func (u *User) ActivateUser() {
	u.Status = UserStatusActive
}

func (u *User) DeactivateUser() {
	u.Status = UserStatusInactive
}

func (u *User) SetPassword(password string) {
	hashedPassword, _ := util.HashPassword(password)
	u.Password = hashedPassword
}

func (u *User) CheckPassword(password string) bool {
	return util.CheckPassword(password, u.Password)
}

func (u *User) Maskfields() {
	u.Password = "*******"
}
