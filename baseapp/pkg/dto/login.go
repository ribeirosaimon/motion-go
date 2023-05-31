package dto

import "github.com/ribeirosaimon/motion-go/internal/domain/sqlDomain"

type SignUpDto struct {
	LoginDto
	Name  string               `json:"name"`
	Roles []sqlDomain.RoleEnum `json:"roles"`
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}
