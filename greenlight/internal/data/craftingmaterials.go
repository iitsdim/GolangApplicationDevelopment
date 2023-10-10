package data

import (
	"time"
)

type CraftingMaterials struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Price     Price     `json:"price,string"`
	CreatedAt time.Time `json:"-"`
}
