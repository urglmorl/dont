package models

import (
	"github.com/google/uuid"
)

type Theme struct {
	Id        uuid.UUID
	Name      string
	Questions []Question
	SubjectId uuid.UUID
}
