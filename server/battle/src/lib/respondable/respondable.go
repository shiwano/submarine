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
func (t *T) Respond(result interface{}, err error) {
	t.done <- response{
		value: result,
		err:   err,
	}
}

// Receive a value from the channel.
func (t *T) Receive() (interface{}, error) {
	response := <-t.done
	close(t.done)
	return response.value, response.err
}

type response struct {
	value interface{}
	err   error
}
