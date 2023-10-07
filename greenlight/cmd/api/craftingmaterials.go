package main

import (
	"fmt"
	"greenlight.dimash.net/internal/data"
	"net/http"
	"time"
)

// Add a createMovieHandler for the "POST /v1/crafting_materials" endpoint. For now we simply
// return a plain-text placeholder response.
func (app *application) createCraftingMaterialHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new crafting Material")
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
	}

	err = app.writeJSON(w, http.StatusOK, craftingMaterial, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered "+
			"a problem and could not process your request", http.StatusInternalServerError)
	}
}
