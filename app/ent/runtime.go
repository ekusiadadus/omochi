// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/YadaYuki/omochi/app/ent/schema"
	"github.com/YadaYuki/omochi/app/ent/term"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	termFields := schema.Term{}.Fields()
	_ = termFields
	// termDescCreatedAt is the schema descriptor for created_at field.
	termDescCreatedAt := termFields[2].Descriptor()
	// term.DefaultCreatedAt holds the default value on creation for the created_at field.
	term.DefaultCreatedAt = termDescCreatedAt.Default.(time.Time)
	// termDescUpdatedAt is the schema descriptor for updated_at field.
	termDescUpdatedAt := termFields[3].Descriptor()
	// term.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	term.DefaultUpdatedAt = termDescUpdatedAt.Default.(time.Time)
	// termDescID is the schema descriptor for id field.
	termDescID := termFields[0].Descriptor()
	// term.DefaultID holds the default value on creation for the id field.
	term.DefaultID = termDescID.Default.(func() uuid.UUID)
}
