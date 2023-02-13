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
	{
		URI:      "/user/follower/{id}",
		Method:   http.MethodGet,
		Function: controllers.AddFollowings,
		Auth:     false,
	},
	{
		URI:      "/user/login",
		Method:   http.MethodPost,
		Function: controllers.Login,
		Auth:     false,
	},
}
