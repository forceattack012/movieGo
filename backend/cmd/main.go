package main

import (
	"backend/internal/database/mongo"
	"backend/internal/database/postgre"
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/routes"
	"backend/internal/services"
	"fmt"
	"log"
	"net/http"
	"time"

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

	if err != nil {
		log.Println(err)
		return
	}

	hostPg := viper.GetString("postgrsdb.host")
	user := viper.GetString("postgrsdb.user")
	password := viper.GetString("postgrsdb.password")
	db := viper.GetString("postgrsdb.databaseName")
	portPg := viper.GetString("postgrsdb.port")

	configPg := postgre.BuildConfig(hostPg, user, password, db, portPg)
	postgre.Connect(configPg)

	postgre.Database.AutoMigrate(&models.Theater{}, &models.ShowTime{}, &models.Seat{})

	movieRepo := repositories.NewMovieRepository(mongoDb, "Movies")
	movieService := services.NewMovieService(movieRepo)

	theaterRepo := repositories.NewTheaterRepository(postgre.Database)
	theaterService := services.NewTheaterService(theaterRepo)

	showTimeRepo := repositories.NewShowTimeRepository(postgre.Database)
	showTimeService := services.NewShowTimeService(showTimeRepo, movieRepo)

	e := echo.New()
	e.Validator = &ValidateRequest{
		v: validator.New(),
	}

	routes.HelloRoutes(e)
	routes.NewMovieRouter(e, movieService)
	routes.NewTheaterRouter(e, theaterService)
	routes.NewShowTimeRoute(e, showTimeService)

	loc, err := time.LoadLocation("Asia/Bangkok")
	time.Local = loc

	e.Logger.Fatal(e.Start(":5001"))
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
