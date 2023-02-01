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
}
