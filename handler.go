package main

import (
	"encoding/json"
	"net/http"
)

type HelloHandler struct {
	Saying  string
	counter int
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	saying := Saying{
		Id:     h.counter,
		Saying: h.Saying,
	}

	b, _ := json.Marshal(&saying)
	h.counter++
	w.Write(b)
}
