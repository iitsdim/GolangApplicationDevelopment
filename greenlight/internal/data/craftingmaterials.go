package data

import (
	"greenlight.dimash.net/internal/validator"
	"time"
)

type CraftingMaterials struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Price     Price     `json:"price,string"`
	CreatedAt time.Time `json:"-"`
}

func ValidateMovie(v *validator.Validator, materials CraftingMaterials) {
	v.Check(materials.Title != "", "title", "must be provided")
	v.Check(len(materials.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(materials.Year >= 1888, "year", "must be greater than 1888")
	v.Check(materials.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	v.Check(materials.Price > 0, "runtime", "must be a positive integer")
}
