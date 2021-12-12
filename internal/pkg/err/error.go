package err

type Error struct {
	message string
}

func (e Error) Error() string {
	return e.message
}
func NewError(message string) error {
	return &Error{message: message}
}
