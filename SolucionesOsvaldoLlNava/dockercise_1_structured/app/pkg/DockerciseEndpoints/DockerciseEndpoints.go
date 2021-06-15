package DockerciseEndpoints

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/OsvaldoLlNava/dockercises1/structured/app/pkg/DockerciseDatabase"
	"github.com/OsvaldoLlNava/dockercises1/structured/app/pkg/DockerciseModels"
	"github.com/go-chi/chi/v5"
)

func AllResults(w http.ResponseWriter, r *http.Request) {
	p := obtenerTodo()

	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		panic(err)
	}
}

func SpecificResult(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "userId")
	p := obtenerPersona(id)
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		panic(err)
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./app/resources/templates/home.html")
	if err != nil {
		fmt.Fprintf(w, "Bienvenido al Home")
	}

	err = t.Execute(w, nil)
	if err != nil {
		fmt.Fprintf(w, "Pagina de home temporalmente fuera de servicio")
	}
}

func obtenerTodo() []DockerciseModels.Person {

	p := DockerciseDatabase.ObtenerPersonas()

	return p
}

func obtenerPersona(i string) DockerciseModels.Person {
	p := DockerciseDatabase.ObtenerUnaPersona(i)
	return p
}
