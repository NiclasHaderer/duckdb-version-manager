package stacktrace

type Error interface {
	error
	StackTrace() string
}

type trackableError struct {
	error error
	stack string
}

func (t trackableError) Error() string {
	return t.error.Error()
}

func (t trackableError) StackTrace() string {
	return t.stack
}
