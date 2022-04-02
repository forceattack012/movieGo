package main

import (
	"backend/internal/database/mongo"
	"backend/internal/repositories"
	"backend/internal/routes"
	"backend/internal/services"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

var err error

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

	log.Println(viper.GetString("app.env"))

	host := viper.GetString("mongodb.host")
	port := viper.GetString("mongodb.port")
	databaseName := viper.GetString("mongodb.databaseName")

	config := mongo.BuildConfig(host, port, databaseName)
	mongoDb, err := mongo.Connect(config)
	mongo.DeleteSeed(mongoDb, "Movies")
	mongo.Seed(mongoDb, "Movies")

	if err != nil {
		log.Println(err)
		return
	}

	// configPg := postgre.BuildConfig()
	// postgre.Database, err = postgre.Connect(configPg)

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// log.Println(postgre.Database)

	movieRepo := repositories.NewMovieRepository(mongoDb, "Movies")
	movieService := services.NewMovieService(movieRepo)

	theaterRepo := repositories.NewTheaterRepository(mongoDb, "Theater")
	theaterService := services.NewTheaterService(theaterRepo)

	e := echo.New()
	e.Validator = &ValidateRequest{
		v: validator.New(),
	}

	routes.HelloRoutes(e)
	routes.NewMovieRouter(e, movieService)
	routes.NewTheaterRouter(e, theaterService)

	e.Logger.Fatal(e.Start(":3000"))
}

type ValidateRequest struct {
	v *validator.Validate
}

func (cv *ValidateRequest) Validate(i interface{}) error {
	if err := cv.v.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}
