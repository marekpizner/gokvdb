package command

type Response interface {
	// Val() interface{}
}
type ErrResult struct {
	err error
}
type OkResult struct{}

type NilReply struct{}

//Val returns nil.
func (nr NilReply) Val() interface{} {
	return nil
}

//StringReply contains a string.
type StringReply struct {
	Value string
}

//Val returns string.
func (sr StringReply) Val() interface{} {
	return sr.Value
}
