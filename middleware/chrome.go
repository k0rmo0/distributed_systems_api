package middleware

import (
	"net/http"

	"github.com/ismar/dsa/distrybuted_systems_api/utils"
)

//Middleware ...
type Middleware struct{}

// CORS ...
func (m Middleware) CORS(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	// CORS support for Preflighted requests
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, PATCH, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

	next(res, req)
}

// Preflight ...
func (m Middleware) Preflight(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	if req.Method == "OPTIONS" {
		utils.Renderer.Render(res, http.StatusOK, map[string]string{"status": "OK"})
		return
	}

	next(res, req)
}
