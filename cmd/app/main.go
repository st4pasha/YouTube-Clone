package main

import (
	"net/http"

	"github.com/st4pasha/YouTube-Clone/cmd/internal/handler"
)

func main(){
	router := handler.RegisterHandlers()

	if err :=  http.ListenAndServe(":8080", router); err != nil {
		return
	}
}