package databases

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type GormJSON map[string]any

func (j GormJSON) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}

func (j *GormJSON) Scan(value any) error {
	if value == nil {
		*j = make(GormJSON)
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("invalid type for GormJSON")
	}

	return json.Unmarshal(bytes, &j)
}
