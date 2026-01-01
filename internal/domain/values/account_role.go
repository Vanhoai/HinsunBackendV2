package values

import (
	"fmt"
	"hinsun-backend/internal/core/failure"
)

type AccountRole int

const (
	NormalRole AccountRole = iota
	AdminRole
	GodRole
)

func RoleFromInt(role int) (AccountRole, error) {
	switch role {
	case 0:
		return NormalRole, nil
	case 1:
		return AdminRole, nil
	case 2:
		return GodRole, nil
	default:
		return -1, failure.NewValidationFailure(fmt.Sprintf("Invalid account role: %d", role))
	}
}
