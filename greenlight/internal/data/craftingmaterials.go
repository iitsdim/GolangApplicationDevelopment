package data

import (
	"time"
)

type CraftingMaterials struct {
	ID        int64
	Title     string
	Year      int32
	CreatedAt time.Time
}
