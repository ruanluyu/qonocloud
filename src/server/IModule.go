package server

import (
	"net/http"
)

type ModuleContext struct{
	Request			*http.Request
	Response		http.ResponseWriter
	Fallback 		func()
}

type IModule interface {
	Run(context *ModuleContext) error
}
