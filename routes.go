package main

import (
	"net/http"
	"strength-app/handlers"

	"github.com/jmoiron/sqlx"
)

type Route struct {
	Path    string
	Method  string
	Handler func(*http.Request, map[string]string, *sqlx.DB) (handlers.Response, error)
}

func Routes() []Route {
	return []Route{
		// Exercises
		{"/exercises", "GET", handlers.Exercise{}.Index},
		{"/exercises/{id}", "GET", handlers.Exercise{}.Show},
		{"/exercises", "POST", handlers.Exercise{}.Store},
		{"/exercises/{id}", "PUT", handlers.Exercise{}.Update},
		{"/exercises/{id}", "DELETE", handlers.Exercise{}.Destroy},

		// Macrocycles
		{"/macrocycles", "GET", handlers.Macrocycle{}.Index},
		{"/macrocycles/{id}", "GET", handlers.Macrocycle{}.Show},
		{"/macrocycles", "POST", handlers.Macrocycle{}.Store},
		{"/macrocycles/{id}", "PUT", handlers.Macrocycle{}.Update},
		{"/macrocycles/{id}", "DELETE", handlers.Macrocycle{}.Destroy},

		// Mesocycles
		{"/mesocycles", "GET", handlers.Mesocycle{}.Index},
		{"/mesocycles/{id}", "GET", handlers.Mesocycle{}.Show},
		{"/mesocycles", "POST", handlers.Mesocycle{}.Store},
		{"/mesocycles/{id}", "PUT", handlers.Mesocycle{}.Update},
		{"/mesocycles/{id}", "DELETE", handlers.Mesocycle{}.Destroy},

		// Microcycles
		{"/microcycles", "GET", handlers.Microcycle{}.Index},
		{"/microcycles/{id}", "GET", handlers.Microcycle{}.Show},
		{"/microcycles", "POST", handlers.Microcycle{}.Store},
		{"/microcycles/{id}", "PUT", handlers.Microcycle{}.Update},
		{"/microcycles/{id}", "DELETE", handlers.Microcycle{}.Destroy},

		// Sessions
		{"/sessions", "GET", handlers.Session{}.Index},
		{"/sessions/{id}", "GET", handlers.Session{}.Show},
		{"/sessions", "POST", handlers.Session{}.Store},
		{"/sessions/{id}", "PUT", handlers.Session{}.Update},
		{"/sessions/{id}", "DELETE", handlers.Session{}.Destroy},

		// Slots
		{"/slots", "GET", handlers.Slot{}.Index},
		{"/slots/{id}", "GET", handlers.Slot{}.Show},
		{"/slots", "POST", handlers.Slot{}.Store},
		{"/slots/{id}", "PUT", handlers.Slot{}.Update},
		{"/slots/{id}", "DELETE", handlers.Slot{}.Destroy},

		// Sets
		{"/sets", "GET", handlers.Set{}.Index},
		{"/sets/{id}", "GET", handlers.Set{}.Show},
		{"/sets", "POST", handlers.Set{}.Store},
		{"/sets/{id}", "PUT", handlers.Set{}.Update},
		{"/sets/{id}", "DELETE", handlers.Set{}.Destroy},
	}
}
