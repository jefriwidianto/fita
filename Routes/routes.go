package Routes

import (
	"fita/Controller"
	"github.com/go-chi/chi"
	"net/http"
)

type Routes struct {
	Controller Controller.InterfaceController
	Chi        *chi.Mux
}

func (app *Routes) CollectRoutes() {
	appRoutes := chi.NewRouter()
	appRoutes.Post("/transaction", app.Controller.Transaction)


	http.ListenAndServe("localhost:3000", appRoutes)
}
