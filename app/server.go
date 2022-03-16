package main

import (
	"database/sql"
	// "encoding/json"
	// "os"
	// "io/ioutil"
	"log"
    "net/http"
    "github.com/labstack/echo/v4"
	"fmt"
	_ "github.com/lib/pq"
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

	var Db *sql.DB
	Db, err := sql.Open("postgres", "host=postgres port=5432 user=app_user password=app_password dbname=app_db sslmode=disable")
    if err != nil {
		log.Fatal(err)
	}

	// sql := "SELECT id, station_name, line_name FROM STATION_MATSTER_TB WHERE id=$1;"

	// pstatement, err := Db.Prepare(sql)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// queryID := 1
	// var testStation Station

	// err = pstatement.QueryRow(queryID).Scan(&testStation.Id, &testStation.StationName, &testStation.LineName)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(testStation.Id, testStation.StationName, testStation.LineName)

	e := echo.New()
    e.GET("/getStations", func(c echo.Context) error {
		station := new(Station)
		var stations []Station

		// stationsJsonFile, err := os.Open("./TestPackages/json/stations.json")
	
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// defer stationsJsonFile.Close()

		// stationsByteValue, _ := ioutil.ReadAll(stationsJsonFile)
		// json.Unmarshal(stationsByteValue, &stations)
		
		sql := "SELECT id, station_name, line_name FROM STATION_MATSTER_TB;"

		rows, err := Db.Query(sql)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next(){
			if err := rows.Scan(&station.Id, &station.StationName, &station.LineName); err != nil {
				log.Fatal(err)
			}
			stations = append(stations, Station{
					Id: station.Id,
					StationName: station.StationName,
					LineName: station.LineName,
				})
		}

		response := new(Response)

		response.Stations = stations

		fmt.Println(stations)
		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "http://localhost:4200")
        return c.JSON(http.StatusOK, response)
    })
    e.Logger.Fatal(e.Start(":8080"))
}