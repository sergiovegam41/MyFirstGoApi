package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := gin.Default()

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://stock-manager:r4mEHcjNNw3z9u3K@cluster0.0eyr1.mongodb.net/stock-manager?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	// Use a longer timeout for initial connection.
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err = client.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(nil)

	database := client.Database("test-dservice-backend")
	collection := database.Collection("departments")

	r.GET("/getCountries", func(c *gin.Context) {
		// Create a new context for this request.
		reqCtx, reqCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer reqCancel()

		var results []bson.M

		cursor, err := collection.Find(reqCtx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al consultar MongoDB", "error": err.Error()})
			return
		}
		defer cursor.Close(reqCtx)

		if err = cursor.All(reqCtx, &results); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al leer documentos", "error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, results)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
