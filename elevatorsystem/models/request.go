package models

type Request struct {
	sourceFloor      int
	destinationFloor int
}

func NewRequest(src, dst int) *Request {
	return &Request{
		sourceFloor:      src,
		destinationFloor: dst,
	}
}
