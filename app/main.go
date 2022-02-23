package main

import (
	"encoding/json"
	"os"
	"net"
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
	
	var station Station
	fmt.Println(station)
	jsonErr := json.Unmarshal(stationsByteValue, &station)
	fmt.Println(station)
	fmt.Println(reflect.TypeOf(station))
	log.Fatal(jsonErr)
	return c.JSON(http.StatusOK, station)
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:8080")
	defer l.Close()
	if err !=nil {
		panic(err)
	}
	fmt.Println("wait...")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})
	e.GET("/getStations", getStations)
	e.Logger.Fatal(e.Start(":8080"))
}