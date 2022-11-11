package routes

import (
	"api/src/controllers"
	"net/http"
)

var UserRoutes = []Route{
	{
		Uri:         "/users",
		Method:      http.MethodPost,
		Function:    controllers.CraeteUser,
		RequestAuth: false,
	},
	{
		Uri:         "/users",
		Method:      http.MethodGet,
		Function:    controllers.GrabAllUsers,
		RequestAuth: false,
	},
	{
		Uri:         "/users/{id}",
		Method:      http.MethodGet,
		Function:    controllers.GrabUser,
		RequestAuth: false,
	},
	{
		Uri:         "/users/{id}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequestAuth: false,
	},
	{
		Uri:         "/users/{id}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequestAuth: false,
	},
}
