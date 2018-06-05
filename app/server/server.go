package server

import (
	"fmt"
	"net/http"

	"github.com/ttakezawa/go-service-example/usecase"
)

// App is web server.
type App struct {
	UserUsecase *usecase.UserUsecase `inject:""`
}

// Run as a server.
func (app *App) Run() error {
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[7:]
		user, err := app.UserUsecase.Get(r.Context(), name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error: %s", err.Error())
			return
		}
		if user == nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "user not found")
			return
		}
		fmt.Fprintf(w, "user: %v", user)
	})
	return http.ListenAndServe(":3000", nil)
}
