package main

import (
	"encoding/json"
	"os"
	"io/ioutil"
	"log"
    "net/http"
    "github.com/labstack/echo/v4"
	"fmt"
)

type Station struct {
	Id int `json:"id"`
	StationName string `json:"stationName"`
	LineName string `json:"lineName"`
}

type Response struct {
	Stations []Station `json:"stations"`
}

func main() {
    e := echo.New()
    e.GET("/getStations", func(c echo.Context) error {
		var stations []Station

		stationsJsonFile, err := os.Open("./TestPackages/json/stations.json")
	
		if err != nil {
			log.Fatal(err)
		}

		defer stationsJsonFile.Close()

		stationsByteValue, _ := ioutil.ReadAll(stationsJsonFile)
		json.Unmarshal(stationsByteValue, &stations)
		response := new(Response)
		response.Stations = stations
		fmt.Println(stations)
        return c.JSON(http.StatusOK, response)
    })
    e.Logger.Fatal(e.Start(":8080"))
}