package store

import (
	"github.com/Max425/film-library.git/internal/domain"
	"time"
)

type User struct {
	ID           int       `db:"id"`
	Name         string    `db:"name"`
	Mail         string    `db:"mail"`
	PasswordHash string    `db:"password_hash"`
	Salt         string    `db:"salt"`
	Role         int       `db:"role"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func UserDomainToStore(domainUser *domain.User) *User {
	return &User{
		ID:           domainUser.ID(),
		Name:         domainUser.Name(),
		Mail:         domainUser.Mail(),
		PasswordHash: domainUser.Password(),
		Salt:         domainUser.Salt(),
		Role:         domainUser.Role(),
	}
}

func UserStoreToDomain(storeUser *User) (*domain.User, error) {
	return domain.NewUser(
		storeUser.ID,
		storeUser.Name,
		storeUser.Mail,
		storeUser.PasswordHash,
		storeUser.Salt,
		storeUser.Role,
	)
}
