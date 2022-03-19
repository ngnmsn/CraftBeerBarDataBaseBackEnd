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
	StationId int `json:"stationId"`
	StationName string `json:"stationName"`
	LineName string `json:"lineName"`
}

type GetStationsResponse struct {
	Stations []Station `json:"stations"`
}

type Bar struct {
	BarId int `json:"barId"`
	BarName string `json:"barName"`
	NumberOfTaps int `json:"numberOfTaps"`
	Address string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	TabelogLink string `json:"tabelogLink"`	
}

type GetBarListResponse struct {
	BarList []Bar `json:"barList"`
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
		
		sql := "SELECT station_id, station_name, line_name FROM STATION_MATSTER_TB;"

		rows, err := Db.Query(sql)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next(){
			if err := rows.Scan(&station.StationId, &station.StationName, &station.LineName); err != nil {
				log.Fatal(err)
			}
			stations = append(stations, Station{
					StationId: station.StationId,
					StationName: station.StationName,
					LineName: station.LineName,
				})
		}

		getStationsResponse := new(GetStationsResponse)

		getStationsResponse.Stations = stations

		fmt.Println(stations)
		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "http://localhost:4200")
        return c.JSON(http.StatusOK, getStationsResponse)
    })
	e.GET("/getBarList", func(c echo.Context) error {

		bar := new(Bar)
		var barList []Bar

		stationId := c.QueryParam("stationId")
		fmt.Println(stationId)

		sql := "SELECT a.bar_id, a.bar_name, a.number_of_taps, a.address, a.phone_number, a.tabelog_link FROM BAR_MASTER_TB a, NEAREST_STATION_TB b WHERE b.station_id = $1 AND b.bar_id = a.bar_id;"

		rows, err := Db.Query(sql, stationId)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next(){
			if err := rows.Scan(&bar.BarId, &bar.BarName, &bar.NumberOfTaps, &bar.Address, &bar.PhoneNumber, &bar.TabelogLink); err != nil {
				log.Fatal(err)
			}
			barList = append(barList, Bar{
					BarId: bar.BarId,
					BarName: bar.BarName,
					NumberOfTaps: bar.NumberOfTaps,
					Address: bar.Address,
					PhoneNumber: bar.PhoneNumber,
					TabelogLink: bar.TabelogLink,
				})
		}

		getBarListResponse := new(GetBarListResponse)

		getBarListResponse.BarList = barList

		fmt.Println(barList)
		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "http://localhost:4200")
        return c.JSON(http.StatusOK, getBarListResponse)
    })
    e.Logger.Fatal(e.Start(":8080"))
}