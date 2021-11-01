package generator

import "github.com/google/uuid"

// GenerateUUID generates a unique ID that can be used as an identifier for an entity or else.
func GenerateUUID() string {
	return uuid.New().String()
}
