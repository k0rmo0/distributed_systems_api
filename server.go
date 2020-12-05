package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ismar/dsa/distrybuted_systems_api/controllers"
	"github.com/ismar/dsa/distrybuted_systems_api/utils"
	"github.com/julienschmidt/httprouter"
)

var (
	val controllers.MeasuredValues
)

var err error

func main() {
	//Making initial DB Access
	utils.GetSQLDB()

	//Router
	mux := httprouter.New()

	mux.POST("/measured-distance", val.Save)
	mux.GET("/data", val.GetData)

	fmt.Println("Server started...")
	log.Fatal(http.ListenAndServe(":8080", mux))

}
