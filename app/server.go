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
	OfficialLink string `json:"officialLink"`
	TabelogLink string `json:"tabelogLink"`
	Food string `json:"food"`
	Style string `json:"style"`
	NearStationName string `json:"nearStationName"`
	OnFootTime int `json:"onFootTime"`
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

		keyword := c.QueryParam("keyword")
		fmt.Println(keyword)
		search := "%"
		fmt.Println(search)

		keywordSearch := keyword + search

		sql := "SELECT a.bar_id, a.bar_name, a.number_of_taps, a.address, a.phone_number, a.official_link, a.tabelog_link, a.food, a.style, c.station_name, b.on_foot_time FROM BAR_MASTER_TB a, NEAREST_STATION_TB b, STATION_MATSTER_TB c WHERE c.station_name LIKE $1 AND c.station_id = b.station_id AND b.bar_id = a.bar_id;"

		rows, err := Db.Query(sql, keywordSearch)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next(){
			if err := rows.Scan(&bar.BarId, &bar.BarName, &bar.NumberOfTaps, &bar.Address, &bar.PhoneNumber, &bar.OfficialLink, &bar.TabelogLink, &bar.Food, &bar.Style, &bar.NearStationName, &bar.OnFootTime); err != nil {
				log.Fatal(err)
			}
			barList = append(barList, Bar{
					BarId: bar.BarId,
					BarName: bar.BarName,
					NumberOfTaps: bar.NumberOfTaps,
					Address: bar.Address,
					PhoneNumber: bar.PhoneNumber,
					OfficialLink: bar.OfficialLink,
					TabelogLink: bar.TabelogLink,
					Food: bar.Food,
					Style: bar.Style,
					NearStationName: bar.NearStationName,
					OnFootTime: bar.OnFootTime,
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