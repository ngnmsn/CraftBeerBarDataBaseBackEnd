package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

type ItemData struct {
    Id int `json:"id"`
    Name string `json:"name"`
}

func main() {
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        item := new(ItemData)
        if err := c.Bind(item); err != nil {
            return err
        }
        return c.JSON(http.StatusOK, item)
    })
    e.Logger.Fatal(e.Start(":8080"))
}