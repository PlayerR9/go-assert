package assert

// ErrAssertFailed is an error returned when an assertion fails.
type ErrAssertFailed struct {
	// Msg is the error message
	Msg string
}

// Error implements the error interface.
//
// Message:
//
//	"[ASSERT FAILED]: <msg>"
func (e ErrAssertFailed) Error() string {
	msg := "[ASSERT FAILED]: "

	if e.Msg != "" {
		msg += e.Msg
	} else {
		msg += "something went wrong"
	}

	return msg
}

// NewErrAssertFailed creates a new ErrAssertFailed error.
//
// Parameters:
//   - msg: the error message.
//
// Returns:
//   - *ErrAssertFailed: the new error.
func NewErrAssertFailed(msg string) *ErrAssertFailed {
	return &ErrAssertFailed{
		Msg: msg,
	}
}
