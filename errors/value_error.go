package errors

// ValueError represents the error related to value has been occured.
type ValueError struct {
	Message string
	Value   interface{}
}

func (err *ValueError) Error() string {
	return err.Message
}

// Init creates a new Value error instance
func (err ValueError) Init(message string, Value interface{}) *ValueError {
	err.Message = message
	err.Value = Value
	return &err
}
