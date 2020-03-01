package models

import (
	"github.com/google/uuid"
)

type Option struct {
	Id      uuid.UUID
	Name    string
	IsRight bool
}
