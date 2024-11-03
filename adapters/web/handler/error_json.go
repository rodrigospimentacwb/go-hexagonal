package handler

import "encoding/json"

func jsonError(msg string) []byte {
	msgToParse := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}
	r, err := json.Marshal(msgToParse)
	if err != nil {
		return []byte(err.Error())
	}
	return r
}
