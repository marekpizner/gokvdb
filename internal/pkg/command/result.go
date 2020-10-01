package command

type Result interface {
}
type ErrResult struct {
	err error
}
type OkResult struct{}
