package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SendSuccessResponse sends a success response with status code 200 and JSON body.
func SendSuccessResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// SendErrorResponse sends an error response with the specified status code and message.
func SendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{"error": message})
}

func UnmarshalBody(r *http.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		var msg string
		switch t := err.(type) {
		case *json.SyntaxError:
			jsn := make([]byte, t.Offset)
			r.Body.Read(jsn)
			jsn = append(jsn, '<', '-', '-', '(', 'I', 'n', 'v', 'a', 'l', 'i', 'd', ' ', 'C', 'h', 'a', 'r', 'a', 'c', 't', 'e', 'r', ')')
			msg = fmt.Sprintf("Invalid character at offset %v\n %s", t.Offset, jsn)
		case *json.UnmarshalTypeError:
			jsn := make([]byte, t.Offset)
			r.Body.Read(jsn)
			jsn = append(jsn, '<', '-', '-', '(', 'I', 'n', 'v', 'a', 'l', 'i', 'd', ' ', 'T', 'y', 'p', 'e', ')')
			msg = fmt.Sprintf("Invalid value at offset %v\n %s", t.Offset, jsn)
		default:
			msg = err.Error()
		}

		return fmt.Errorf("400: %s", msg)
	}

	return nil
}
