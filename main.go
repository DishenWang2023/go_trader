package main

import (
	"context"
	"log"
	"time"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func main() {
	c := polygon.New("MtL1p6hyMHWYY7WxCjUPpjZfQkOJFpvh")

	params := models.GetTickerDetailsParams{
		Ticker: "AAPL",
	}.WithDate(models.Date(time.Date(2021, 7, 22, 0, 0, 0, 0, time.Local)))

	res, err := c.GetTickerDetails(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res) // do something with the result
}
