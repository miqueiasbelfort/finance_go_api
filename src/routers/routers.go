package routers

import (
	"api/src/routers/router"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	r := mux.NewRouter()
	return router.ConfigRouter(r)
}
