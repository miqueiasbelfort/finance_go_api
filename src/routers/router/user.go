package router

import (
	"api/src/controllers"
	"net/http"
)

var userRouters = []Route{
	{
		URI:      "/",
		Method:   http.MethodGet,
		Function: controllers.CreateAUser,
		Auth:     false,
	},
}
