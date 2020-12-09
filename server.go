package main

import (
	"fmt"
	"time"

	"github.com/ismar/dsa/distrybuted_systems_api/controllers"
	middleware "github.com/ismar/dsa/distrybuted_systems_api/middleware"
	"github.com/ismar/dsa/distrybuted_systems_api/utils"
	"github.com/julienschmidt/httprouter"
	"github.com/tylerb/graceful"
	"github.com/urfave/negroni"
)

var (
	val  controllers.MeasuredValues
	mdlw middleware.Middleware
)

var err error

func main() {
	//Making initial DB Access
	utils.GetSQLDB()

	n := negroni.Classic()

	n.Use(negroni.HandlerFunc(mdlw.CORS))
	n.Use(negroni.HandlerFunc(mdlw.Preflight))

	port := "8010"
	serverIP := "192.168.0.19"

	recovery := negroni.NewRecovery()
	recovery.PrintStack = false

	n.Use(recovery)
	//Router
	mux := httprouter.New()

	mux.POST("/measured-distance", val.Save)
	mux.GET("/data", val.GetData)

	n.UseHandler(mux)

	fmt.Println("Server started...")
	n.Run(serverIP + ":" + port)

	graceful.Run(serverIP+":"+port, 10*time.Second, n)

}
