package routers

const (
	Error        = "Error"
	ErrorMessage = "Internal server error"
	Insert       = "Se Guardo Correctamente"
	Message      = "message"
)

type response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func NewResponse(messageType string, message string, data interface{}) response {
	return response{
		messageType,
		message,
		data,
	}
}
