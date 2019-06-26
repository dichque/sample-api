// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/dichque/sample-api/models"
	"github.com/go-openapi/swag"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/dichque/sample-api/restapi/operations"
	"github.com/dichque/sample-api/restapi/operations/pet"
)

//go:generate swagger generate server --target ../../sample-api --name MinimalPetStoreExample --spec ../swagger.yaml

var petList = []*models.Pet{
	{ID: 0, Name: swag.String("Bobby"), Kind: "dog"},
	{ID: 1, Name: swag.String("Lola"), Kind: "cat"},
	{ID: 2, Name: swag.String("Bella"), Kind: "dog"},
	{ID: 3, Name: swag.String("Maggie"), Kind: "cat"},
}

func configureFlags(api *operations.MinimalPetStoreExampleAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MinimalPetStoreExampleAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.PetCreateHandler == nil {
		api.PetCreateHandler = pet.CreateHandlerFunc(func(params pet.CreateParams) middleware.Responder {
			return middleware.NotImplemented("operation pet.Create has not yet been implemented")
		})
	}
	if api.PetGetHandler == nil {
		api.PetGetHandler = pet.GetHandlerFunc(func(params pet.GetParams) middleware.Responder {
			return middleware.NotImplemented("operation pet.Get has not yet been implemented")
		})
	}

	api.PetListHandler = pet.ListHandlerFunc(func(params pet.ListParams) middleware.Responder {
		if params.Kind == nil {
			return pet.NewListOK().WithPayload(petList)
		}
		var pets []*models.Pet
		for _, pet := range petList {
			if *params.Kind == pet.Kind {
				pets = append(pets, pet)
			}
		}
		return pet.NewListOK().WithPayload(pets)
	})


	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
