package model

import (
	"github.com/google/uuid"
	"time"
)

// UUID in a separate interface allows us to inject any uuid implementation.
// This is already used for testing.
type UUID interface {
	New() uuid.UUID
}

// Time in a separate interface allows us to inject any time implementation.
// This is already used for testing.
type Time interface {
	Now() time.Time
}
