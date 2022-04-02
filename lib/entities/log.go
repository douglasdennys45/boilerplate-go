package entities

import (
	ex "clean-architeture/lib/entities/errors"
	"clean-architeture/lib/entities/validators"
	"clean-architeture/lib/ports/usecases"
	"encoding/json"
	"time"
)

type Log struct {
	Id          string      `json:"id"`
	Type        string      `json:"type"`
	Description interface{} `json:"description"`
	Service     string      `json:"service"`
	Version     float64     `json:"version"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

func Create(data []byte) (*Log, *usecases.Error) {
	var l *Log
	_ = json.Unmarshal(data, &l)
	err := l.isValid()
	return l, &err
}

func (l *Log) isValid() usecases.Error {
	var errors []string
	validator := validators.Validator{}
	if l.Type == "" {
		errors = append(errors, "type parameter is required")
	}
	if !validator.IsString(l.Type) {
		errors = append(errors, "type parameter must be of type string")
	}

	if l.Service == "" {
		errors = append(errors, "service parameter is required")
	}
	if !validator.IsString(l.Service) {
		errors = append(errors, "service parameter must be of type string")
	}

	if l.Version == 0 {
		errors = append(errors, "version parameter is required")
	}
	if !validator.IsDecimal(l.Version) {
		errors = append(errors, "version parameter must be of type decimal")
	}

	return ex.GenerateError("BAD_REQUEST", 400, "BAD REQUEST", errors[0])
}
