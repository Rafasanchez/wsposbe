package functions

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func setErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	}
	return "Unknown error"
}

func GetErrorMsg(er error) []ErrorMsg {
	var msg []ErrorMsg
	var ve validator.ValidationErrors

	if errors.As(er, &ve) {
		msg = make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			msg[i] = ErrorMsg{fe.Field(), setErrorMsg(fe)}
		}
	} else {
		msg = make([]ErrorMsg, 1)
		msg[0] = ErrorMsg{"", er.Error()}
	}

	return msg
}

func GetInfoMsg(code int) string { // Devuelte el texto de los mensajes del sistema
	var msg string
	switch code {
	case 9999:
		msg = "No se ha podido realizar la operación solicitada."
	case 1000:
		msg = "Operación realizada con éxito."
	case 1001:
		msg = "El registro no existe."
	case 1002:
		msg = "No se pudo crear el registro."
	case 1003:
		msg = "No se pudo actualizar el registro."
	default:
		msg = ""
	}

	return msg
}
