package models

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
)

type Link struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Link      string    `json:"link" db:"link"`
	ShortLink string    `json:"short_link" db:"short_link"`
}

// String is not required by pop and may be deleted
func (l Link) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Links is not required by pop and may be deleted
type Links []Link

// String is not required by pop and may be deleted
func (l Links) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Method to generate random slugs
func (l *Link) RandSeq(tx *pop.Connection, n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789$-_.+!*'()")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	// Recursive search if there is existing slugs
	if err := tx.Where("links.short_link = ?", string(b)).First(l); err == nil {
		return l.RandSeq(tx, 6)
	}

	return string(b)
}
