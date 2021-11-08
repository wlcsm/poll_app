package handler

import (
	"encoding/json"
	"net/http"
)

type Return struct {
	Code   int         `json:"code"`
	ErrMsg string      `json:"err_msg"`
	Data   interface{} `json:"data"`
}

func ParseRequest(w *http.Request, d interface{}) error {
	return nil
}

// Generic
func ReturnErr(w http.ResponseWriter, err error, code int) {
	ReturnResp(w, "", err, code)
}

func ReturnResp(w http.ResponseWriter, data interface{}, err error, cod int) {
	code := 0
	err_msg := ""
	if err != nil {
		code = 1
		err_msg = err.Error()
		data = ""
	}

	msg := Return{
		Code:   code,
		ErrMsg: err_msg,
		Data:   data,
	}

	res, err := json.Marshal(msg)
	if err != nil {
		println("what")
	}

	w.WriteHeader(cod)
	w.Write([]byte(res))
}
