package main

import (
	"github.com/dsfsi/covid19za/api/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"os"
)

func main() {
	api := echo.New()
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())
	api.Use(middleware.CORS())
	api.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	latestUpdateController := controllers.NewLatestUpdateController()
	caseController := controllers.NewCaseController()
	hospitalController := controllers.NewHospitalController()

	api.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "COVID 19 data API for South Africa")
	})

	api.GET("/latest-update", latestUpdateController.GetLatestUpdate)
	api.GET("/hospitals/public", hospitalController.GetPublicHospitals)
	api.GET("/hospitals/private", hospitalController.GetPrivateHospitals)
	api.GET("/cases/confirmed", caseController.GetAllConfirmedCases)
	api.GET("/cases/deaths", caseController.GetAllReportedDeaths)
	api.GET("/cases/timeline/tests", caseController.GetTestingTimeline)
	api.GET("/cases/timeline/provincial/cumulative", caseController.GetCumulativeProvincialTimeline)

	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on %s...\n", addr)
	api.Logger.Fatal(api.Start(addr))
}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return ":5000" + port, nil
	}
	return ":" + port, nil
}
