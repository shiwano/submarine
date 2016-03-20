package respondable

// T represents the general-purpose respondable type with thread safe.
type T struct {
	Value interface{}
	done  chan response
}

// New Respondable.
func New(value interface{}) *T {
	return &T{
		Value: value,
		done:  make(chan response, 1),
	}
}

// Respond a value to the channel.
func (r *T) Respond(result interface{}, err error) {
	r.done <- response{
		value: result,
		err:   err,
	}
}

// Receive a value from the channel.
func (r *T) Receive() (interface{}, error) {
	response := <-r.done
	close(r.done)
	return response.value, response.err
}

type response struct {
	value interface{}
	err   error
}
