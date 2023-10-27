package dto

import (
	"github.com/ribeirosaimon/motion-go/config/domain/sqlDomain"
)

type SignUpDto struct {
	LoginDto
	Name  string               `json:"name"`
	Roles []sqlDomain.RoleEnum `json:"roles"`
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type ValidateEmailDto struct {
	Code string `json:"code"`
}
