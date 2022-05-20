package api

import (
	"net/http"

	"Bankuish/api/routes"
)


func NewServer(port string) *http.Server{
	
	routes.Routes()

	return &http.Server{
		Addr: port,

	}
}