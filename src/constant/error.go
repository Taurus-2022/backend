package constant

type StatusError struct {
	statusCode int
	err        error
}

func NewStatusError(statusCode int, err error) StatusError {
	return StatusError{
		statusCode,
		err,
	}
}
func NewOKStatusError() StatusError {
	return StatusError{
		0,
		nil,
	}
}
func (e StatusError) StatusCode() int {
	return e.statusCode
}

func (e StatusError) Error() string {
	return e.err.Error()
}

func (e StatusError) IsNotOK() bool {
	return e.statusCode != ErrorCodeOK
}

func (e StatusError) IsDBError() bool {
	return e.statusCode == ErrorDbInnerError
}
