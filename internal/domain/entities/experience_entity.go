package entities

import (
	"fmt"
	"hinsun-backend/internal/core/failure"
	"time"

	"github.com/google/uuid"
)

const (
	MaxPositionLength   = 100
	MaxCompanyLength    = 100
	MaxLocationLength   = 100
	MaxTechnologys      = 20
	MaxResponsibilities = 5
)

type ExperienceEntity struct {
	ID               uuid.UUID `json:"id"`
	OrderIdx         int8      `json:"orderIdx"`
	Position         string    `json:"position"`
	Company          string    `json:"company"`
	Location         string    `json:"location"`
	Technologies     []string  `json:"technologies"`
	Responsibilities []string  `json:"responsibilities"`
	Period           string    `json:"period"`
	Extra            any       `json:"extra,omitempty"`
	CreatedAt        int64     `json:"createdAt"`
	UpdatedAt        int64     `json:"updatedAt"`
	DeletedAt        *int64    `json:"deletedAt,omitempty"`
}

func NewExperience(
	orderIdx int8,
	position, company, location string,
	technologies, responsibilities []string,
	period string,
) (*ExperienceEntity, error) {
	if err := ValidatePosition(position); err != nil {
		return nil, err
	}

	if err := ValidateCompany(company); err != nil {
		return nil, err
	}

	if err := ValidateLocation(location); err != nil {
		return nil, err
	}

	if err := ValidateTechnologies(technologies); err != nil {
		return nil, err
	}

	if err := ValidateResponsibilities(responsibilities); err != nil {
		return nil, err
	}

	now := time.Now()
	return &ExperienceEntity{
		ID:               uuid.New(),
		OrderIdx:         orderIdx,
		Position:         position,
		Company:          company,
		Location:         location,
		Technologies:     technologies,
		Responsibilities: responsibilities,
		Period:           period,
		Extra:            nil,
		CreatedAt:        now.Unix(),
		UpdatedAt:        now.Unix(),
		DeletedAt:        nil,
	}, nil
}

func ValidatePosition(position string) error {
	if len(position) > MaxPositionLength {
		return failure.NewValidationFailure(
			fmt.Sprintf("position length exceeds maximum of %d characters", MaxPositionLength),
		)
	}

	return nil
}

func ValidateCompany(company string) error {
	if len(company) > MaxCompanyLength {
		return failure.NewValidationFailure(
			fmt.Sprintf("company length exceeds maximum of %d characters", MaxCompanyLength),
		)
	}

	return nil
}

func ValidateLocation(location string) error {
	if len(location) > MaxLocationLength {
		return failure.NewValidationFailure(
			fmt.Sprintf("location length exceeds maximum of %d characters", MaxLocationLength),
		)
	}

	return nil
}

func ValidateTechnologies(technologies []string) error {
	if len(technologies) > MaxTechnologys {
		return failure.NewValidationFailure(
			fmt.Sprintf("number of technologies exceeds maximum of %d", MaxTechnologys),
		)
	}

	return nil
}

func ValidateResponsibilities(responsibilities []string) error {
	if len(responsibilities) > MaxResponsibilities {
		return failure.NewValidationFailure(
			fmt.Sprintf("number of responsibilities exceeds maximum of %d", MaxResponsibilities),
		)
	}

	return nil
}

func (e *ExperienceEntity) Update(
	orderIdx int8,
	position, company, location string,
	technologies, responsibilities []string,
	period string,
) error {
	if err := ValidatePosition(position); err != nil {
		return err
	}

	if err := ValidateCompany(company); err != nil {
		return err
	}

	if err := ValidateLocation(location); err != nil {
		return err
	}

	if err := ValidateTechnologies(technologies); err != nil {
		return err
	}

	if err := ValidateResponsibilities(responsibilities); err != nil {
		return err
	}

	e.OrderIdx = orderIdx
	e.Position = position
	e.Company = company
	e.Location = location
	e.Technologies = technologies
	e.Responsibilities = responsibilities
	e.Period = period
	e.UpdatedAt = time.Now().Unix()

	return nil
}
