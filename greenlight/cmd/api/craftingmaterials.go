package main

import (
	"errors"
	"fmt"
	"greenlight.dimash.net/internal/data"
	"greenlight.dimash.net/internal/validator"
	"net/http"
)

//	for the "POST /v1/crafting_materials" endpoint. For now we simply
//
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

	craftingMaterial := &data.CraftingMaterials{
		Title: input.Title,
		Year:  input.Year,
		Price: input.Price,
	}

	// Initialize a new Validator instance.
	v := validator.New()
	data.ValidateCraftingMaterial(v, craftingMaterial)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.CraftingMaterials.Insert(craftingMaterial)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/crafting_materials/%d", craftingMaterial.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"crafting_material": craftingMaterial}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

//	for the "GET /v1/crafting_materials/:id" endpoint. For now, we retrieve
//
// the interpolated "id" parameter from the current URL and include it in a placeholder
// response.
func (app *application) showCraftingMaterialHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	craftingMaterial, err := app.models.CraftingMaterials.Get(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}

		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"crafting_materials": craftingMaterial}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateCraftingMaterialHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	craftingMaterial, err := app.models.CraftingMaterials.Get(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}

		return
	}

	var input struct {
		Title *string     `json:"title"`
		Year  *int32      `json:"year"`
		Price *data.Price `json:"price"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Year != nil {
		craftingMaterial.Year = *input.Year
	}
	if input.Title != nil {
		craftingMaterial.Title = *input.Title
	}
	if input.Price != nil {
		craftingMaterial.Price = *input.Price
	}

	v := validator.New()
	data.ValidateCraftingMaterial(v, craftingMaterial)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.CraftingMaterials.Update(craftingMaterial)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"crafting_materials": craftingMaterial}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteCraftingMaterialHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.CraftingMaterials.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}

		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "crafting material successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listCraftingMaterialsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string
		data.Filters
	}
	v := validator.New()

	qs := r.URL.Query()

	input.Title = app.readString(qs, "title", "")
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "title", "year", "price", "-id", "-title", "-year", "-price"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	// Call the GetAll() method to retrieve the movies, passing in the various filter
	// parameters.
	crafting_materials, metadata, err := app.models.CraftingMaterials.GetAll(input.Title, input.Filters)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Send a JSON response containing the movie data.
	err = app.writeJSON(w, http.StatusOK, envelope{
		"crafting_materials": crafting_materials,
		"metadata":           metadata,
	}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
