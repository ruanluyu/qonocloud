package server

import (
	"net/http"
)

type ModuleContext struct{
	output *http.ResponseWriter
	input *http.Request
	fallback func()
}

type IModule interface {
	Run(context *ModuleContext) error
}
