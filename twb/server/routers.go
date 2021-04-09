package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"tapi"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"CreateAccount",
		"POST",
		"/v1/accounts",
		tapi.CreateAccount,
	},
	Route{
		"Login",
		"GET",
		"/v1/accounts",
		tapi.CheckAccount,
	},
	Route{
		"GetTeamInfo",
		"GET",
		"/v1/teams/{team_id}",
		tapi.GetTeamInfo,
	},
	Route{
		"CreateTeamInfo",
		"POST",
		"/v1/teams",
		tapi.CreateTeam,
	},
	Route{
		"CreateArticle",
		"POST",
		"/v1/articles",
		tapi.CreateArticle,
	},
	Route{
		"CreateComment",
		"POST",
		"/v1/articles/{article_id}/comments",
		tapi.CreateComment,
	},
	Route{
		"UpdateScore",
		"POST",
		"/v1/articles/{article_id}/score",
		tapi.UpdateScore,
	},
	Route{
		"GetArticle",
		"GET",
		"/v1/articles/{article_id}",
		tapi.GetArticle,
	},
}
