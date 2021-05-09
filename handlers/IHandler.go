package handlers

type IHandler interface {
	Handling()
}

type MainHandler struct {
	IHandler
}
