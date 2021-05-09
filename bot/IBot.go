package bot

type IBot interface {
	Open()
	Close()
	AddHandlerOnce(handler interface{})
	GetSession() interface{}
}

type Bot struct {
	IBot
}
