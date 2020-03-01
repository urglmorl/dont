package models

import (
	"github.com/google/uuid"
)

type Question struct {
	Id        uuid.UUID
	Name      string
	Options   []Option
	SubjectId uuid.UUID
	ThemeId   uuid.UUID
}
