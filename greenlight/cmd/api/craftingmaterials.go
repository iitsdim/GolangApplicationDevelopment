package main

import (
	"fmt"
	"greenlight.dimash.net/internal/data"
	"greenlight.dimash.net/internal/validator"
	"net/http"
	"time"
)

// Add a createMovieHandler for the "POST /v1/crafting_materials" endpoint. For now we simply
// return a plain-text placeholder response.
func (app *application) createCraftingMaterialHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string     `json:"title"`
		Year  int32      `json:"year"`
		Price data.Price `json:"price"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	material := data.CraftingMaterials{
		Title: input.Title,
		Year:  input.Year,
		Price: input.Price,
	}

	// Initialize a new Validator instance.
	v := validator.New()
	data.ValidateMovie(v, material)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

// Add a showMovieHandler for the "GET /v1/crafting_materials/:id" endpoint. For now, we retrieve
// the interpolated "id" parameter from the current URL and include it in a placeholder
// response.
func (app *application) showCraftingMaterialHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	craftingMaterial := data.CraftingMaterials{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Speciality Papers",
		Year:      2022,
		Price:     10000,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"crafting_materials": craftingMaterial}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
