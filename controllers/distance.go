package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ismar/dsa/distrybuted_systems_api/model"
	"github.com/ismar/dsa/distrybuted_systems_api/utils"
	"github.com/julienschmidt/httprouter"
)

// MeasuredValues ...
type MeasuredValues struct{}

// Save ...
func (mv MeasuredValues) Save(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var measuredData model.Values

	queryValues := r.URL.Query()
	measuredData.Time = time.Now().Format("Jan 2 2006 15:04:05")
	distance, _ := strconv.ParseFloat(queryValues.Get("distance"), 64)

	measuredData.Distance = distance

	err := measuredData.WriteToDataBase()

	if err != nil {
		utils.WriteJSON(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, "Succesfully wrote data to database", http.StatusOK)
}

// GetData ...
func (mv MeasuredValues) GetData(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	distanceValues, err := model.ListValues()

	if err != nil {
		fmt.Println(err)
		utils.WriteJSON(w, "Error listing values", http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, distanceValues, http.StatusOK)
}
