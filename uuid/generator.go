// Package uuid provides a pluggable UUID generator interface for consistent ID creation,
// which is especially useful for testing and mocking purposes.
package uuid

import "github.com/google/uuid"

// Generator defines an interface for UUID generation.
type Generator interface {
	// Generate returns a new UUID as a string.
	Generate() string
}

// generator is the default implementation of the Generator interface.
type generator struct{}

// NewGenerator returns a new instance of the default UUID generator.
func NewGenerator() *generator {
	return &generator{}
}

// Generate returns a newly generated UUID string using github.com/google/uuid.
func (g *generator) Generate() string {
	return uuid.New().String()
}
