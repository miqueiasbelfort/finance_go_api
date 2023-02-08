package router

import (
	"api/src/controllers"
	"net/http"
)

var userRouters = []Route{
	{
		URI:      "/user/create",
		Method:   http.MethodPost,
		Function: controllers.CreateAUser,
		Auth:     false,
	},
	{
		URI:      "/user/{id}",
		Method:   http.MethodGet,
		Function: controllers.GetAUser,
		Auth:     false,
	},
	{
		URI:      "/user/{id}",
		Method:   http.MethodPut,
		Function: controllers.UpdateAUser,
		Auth:     false,
	},
}
