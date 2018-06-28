package router

import (
	"server/handler"
)

type Routes []Route

var routes = Routes{
	{"Home", "/", []string{"GET"}, handler.Index, true},
	{"Login", "/login", []string{"POST"}, handler.Login, true},
}
