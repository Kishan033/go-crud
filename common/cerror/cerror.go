package cerror

import "fmt"

var (
	ErrInvalidArgument = New("E001", "invalid argument")
	ErrUnknownError    = func(err error) error {
		return New("E002", err.Error())
	}
	// Business errors
	ErrInvalidPuckPosition      = New("E003", "Invalid Puck Position")
	ErrInvalidPuck              = New("E004", "INVALID_PUCK_ID")
	ErrInvalidPottedPuck        = New("E005", "INVALID_POTTED_PUCK")
	ErrInvalidStrickerPosition  = New("E006", "INVALID_STRIKER_POSITION")
	ErrInvalidStrickerVelocity  = New("E007", "INVALID_STRIKER_POWER")
	ErrInvalidInvalidPlayerTurn = New("E008", "INVALID_PLAYER_TURN")
	ErrInvalidMatchState        = New("E009", "Invalid Match State")
	ErrMatchNotFound            = New("E010", "Match Not Found")
	ErrStateNotFound            = New("E011", "State Not Found")
)

// Error defines specific error that is releated to business error
type Error struct {
	code string
	msg  string
}

// New will return new Error
func New(code string, message string) error {
	return &Error{
		code: code,
		msg:  message,
	}
}

// Error will give string representation of Error
func (e *Error) Error() string { return fmt.Sprintf("Error Code: %s, Error: %s", e.code, e.msg) }

func (e *Error) Code() string {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}
