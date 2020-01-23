package utils

import "github.com/google/uuid"

func UUIDs() uuid.UUID {
	return uuid.New()
}
