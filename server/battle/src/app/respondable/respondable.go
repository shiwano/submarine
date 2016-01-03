package respondable

// Respondable is the general-purpose respondable type for channels.
type Respondable struct {
	Value    interface{}
	response interface{}
	done     chan interface{}
	err      error
}

// New Respondable.
func New(value interface{}) *Respondable {
	return &Respondable{
		Value: value,
		done:  make(chan interface{}, 1),
	}
}

// Respond a value to the channel.
func (r *Respondable) Respond(response interface{}, err error) {
	r.err = err
	r.done <- response
}

// Receive a value from the channel.
func (r *Respondable) Receive() (interface{}, error) {
	r.response = <-r.done
	close(r.done)
	return r.response, r.err
}
