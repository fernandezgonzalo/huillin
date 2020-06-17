package main

import (
	"context"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"fmt"
	log "github.com/sirupsen/logrus"
	"sarlanga/api/handler"
	"sarlanga/db"
	"os"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT")))
	MongoClient, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	dal := db.NewMongoDAL(MongoClient)
	h := &handler.Handler{DAL: dal}

	// LOANS ENDPOINT
	e.GET("/loans/:id", h.GetLoan)
	e.GET("/loans", h.GetLoans)

	// DEBTS ENDPOINT
	e.GET("/debts/:id", h.GetDebt)
	e.GET("/debts", h.GetDebts)

	// INSTALLMENTS ENDPOINT
	e.GET("/installments/:id", h.GetInstallment)
	e.GET("/installments", h.GetInstallments)

	e.Logger.Fatal(e.Start(":1323"))
}

