package server



type IServer interface {
	Init()
	Run()
	Stop()
}
