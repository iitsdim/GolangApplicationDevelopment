package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	// Initialize a new httprouter router instance.
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// Register the relevant methods, URL patterns and handler functions for our
	// endpoints using the HandlerFunc() method

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/crafting_materials", app.listCraftingMaterialsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/crafting_materials", app.createCraftingMaterialHandler)
	router.HandlerFunc(http.MethodGet, "/v1/crafting_materials/:id", app.showCraftingMaterialHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/crafting_materials/:id", app.updateCraftingMaterialHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/crafting_materials/:id", app.deleteCraftingMaterialHandler)

	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}
