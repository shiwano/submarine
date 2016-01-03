package main

// Respondable is the general-purpose respondable type for channel.
type Respondable struct {
	value    interface{}
	response interface{}
	done     chan interface{}
	err      error
}

func newRespondable(value interface{}) *Respondable {
	return &Respondable{value: value, done: make(chan interface{}, 1)}
}

func (r *Respondable) wait() (interface{}, error) {
	r.response = <-r.done
	close(r.done)
	return r.response, r.err
}
