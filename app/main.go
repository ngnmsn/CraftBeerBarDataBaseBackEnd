package main

import (
	"encoding/json"
	"os"
	"net/http"
	"io/ioutil"
	"log"
	"github.com/labstack/echo/v4"
	"fmt"
	"reflect"
)

type Station struct {
	Id int `json:"id"`
	StationName string `json:"stationName"`
	LineName string `json:"lineName"`
}

func getStations(c echo.Context) error {
	stationsJsonFile, err := os.Open("./TestPackages/json/station.json")
	
	if err != nil {
		log.Fatal(err)
	}

	defer stationsJsonFile.Close()

	stationsByteValue, _ := ioutil.ReadAll(stationsJsonFile)
	
	station := new(Station)
	if err := c.Bind(station); err != nil {
		return err
	}
	fmt.Println(station)
	var stations []station
	jsonErr := json.Unmarshal(stationsByteValue, &stations)
	fmt.Println(stations)
	fmt.Println(reflect.TypeOf(stations))
	log.Fatal(jsonErr)
	return c.JSON(http.StatusOK, stations)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})
	e.GET("/getStations", getStations)
	e.Logger.Fatal(e.Start(":8080"))
}