package controllers

import (
	"context"
	"log"
	"net/http"
	"os"

	"tswork-mongo/configs"
	"tswork-mongo/models"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func UploadCsv() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")
		checkError(err)

		err = ctx.SaveUploadedFile(file, "data/"+file.Filename)
		checkError(err)

		var stockdata []models.StockData

		dataFile, err := os.Open("data/" + file.Filename)
		checkError(err)

		err = gocsv.Unmarshal(dataFile, &stockdata)
		checkError(err)
		convertedData := make([]interface{}, len(stockdata))

		for k, v := range stockdata {
			convertedData[k] = v
		}

		var collection *mongo.Collection = configs.GetCollection(configs.DB, "stockData")

		collection.InsertMany(context.Background(), convertedData)

		ctx.JSON(http.StatusOK, gin.H{"message": "data updated successfully"})

	}
}

func GetStockByDate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		date := ctx.Param("date")

		getDate := bson.M{"Date": date}

		var collection *mongo.Collection = configs.GetCollection(configs.DB, "stockData")

		var result models.StockData
		err := collection.FindOne(context.Background(), getDate).Decode(&result)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "stock data for the data not found !"})
			return
		}
		ctx.JSON(http.StatusOK, result)

	}
}
