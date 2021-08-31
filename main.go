package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/pechenini/foodapi/docs"
	"github.com/pechenini/foodapi/model"
	"github.com/pechenini/foodapi/server/handler"
	"github.com/pechenini/foodapi/service/regulation"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	file, err := os.OpenFile("db.yaml", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	dbBytes, err := ioutil.ReadAll(file)
	err = file.Close()
	if err != nil {
		panic(err)
	}

	var db model.DB
	err = yaml.Unmarshal(dbBytes, &db)
	if err != nil {
		panic(err)
	}

	priceRegulator := regulation.NewPriceRegulator(1 * time.Minute)
	go priceRegulator.Regulate(&db)

	restHandler := handler.NewRestaurantHandler(&db)

	r.GET("/restaurants", restHandler.GetRestaurants)

	r.GET("/restaurants/:id/menu", restHandler.GetRestaurantMenu)

	url := ginSwagger.URL(os.Getenv("SWAGGER_ADDRESS")) // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run(os.Getenv("SERVER_ADDRESS")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
