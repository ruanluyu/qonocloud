package server



type IServer interface {
	Init()
	Serve()
	Stop()
}
