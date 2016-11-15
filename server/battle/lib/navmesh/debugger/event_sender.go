package debugger

type eventSender interface {
	Send(event interface{})
}
