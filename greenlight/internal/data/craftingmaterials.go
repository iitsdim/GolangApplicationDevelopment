package data

import (
	"context"
	"database/sql"
	"errors"
	"greenlight.dimash.net/internal/validator"
	"time"
)

type CraftingMaterials struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Price     Price     `json:"price,string"`
	CreatedAt time.Time `json:"-"`
	Version   int32     `json:"version"`
}

func ValidateCraftingMaterial(v *validator.Validator, materials *CraftingMaterials) {
	v.Check(materials.Title != "", "title", "must be provided")
	v.Check(len(materials.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(materials.Year >= 1888, "year", "must be greater than 1888")
	v.Check(materials.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	v.Check(materials.Price > 0, "Price", "must be a positive integer")
}

type CraftingMaterialModel struct {
	DB *sql.DB
}

// Add a placeholder method for fetching a specific record from the movies table.
func (m CraftingMaterialModel) Get(id int64) (*CraftingMaterials, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
	SELECT id, year, price, title, created_at, version 
	from craftingmaterials
	where id = $1`

	var material CraftingMaterials
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&material.ID,
		&material.Year,
		&material.Price,
		&material.Title,
		&material.CreatedAt,
		&material.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &material, nil
}

func (m CraftingMaterialModel) Update(material *CraftingMaterials) error {
	query := `
	UPDATE craftingmaterials
	SET title = $1, year = $2, price = $3, version = version + 1
	WHERE id = $4 AND version = $5
	RETURNING version`

	args := []interface{}{material.Title, material.Year, material.Price, material.ID, material.Version}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&material.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

func (m CraftingMaterialModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
	DELETE from craftingmaterials
	where id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	// Call the RowsAffected() method on the sql.Result object to get the number of rows
	// affected by the query.
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// return an ErrRecordNotFound error.
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (m CraftingMaterialModel) Insert(material *CraftingMaterials) error {
	query := `
	INSERT INTO craftingmaterials (title, year, price) 
	VALUES ($1, $2, $3) RETURNING id, created_at, version`

	args := []interface{}{material.Title, material.Year, material.Price}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	res := m.DB.QueryRowContext(ctx, query, args...).Scan(&material.ID, &material.CreatedAt, &material.Version)
	return res
}

func (m CraftingMaterialModel) GetAll(title string, filters Filters) ([]*CraftingMaterials, error) {
	query := `
	SELECT id, created_at, title, year, price, version 
	from craftingmaterials
	where (LOWER(title) = LOWER($1) OR $1 = '')
	order by id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, title)
	if err != nil {
		return nil, err
	}

	// Importantly, defer a call to rows.Close() to ensure that the resultset is closed
	// before GetAll() returns.

	defer rows.Close()

	// Initialize an empty slice to hold the movie data.
	craftingMaterialsList := []*CraftingMaterials{}

	// Use rows.Next to iterate through the rows in the resultset.
	for rows.Next() {
		// Initialize an empty Movie struct to hold the data for an individual movie.
		var material CraftingMaterials
		// Scan the values from the row into the Crafting Material struct.
		err := rows.Scan(
			&material.ID,
			&material.CreatedAt,
			&material.Title,
			&material.Year,
			&material.Price,
			&material.Version,
		)
		if err != nil {
			return nil, err
		}

		craftingMaterialsList = append(craftingMaterialsList, &material)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return craftingMaterialsList, nil
}
