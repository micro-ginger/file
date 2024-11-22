package handlers

type Request interface {
	WithMaxSize(max uint64) Request
	GetMaxSize() uint64
}

type request struct {
	max uint64
}

func NewRequest() Request {
	return new(request)
}

func (r *request) WithMaxSize(max uint64) Request {
	r.max = max
	return r
}

func (r *request) GetMaxSize() uint64 {
	return r.max
}
