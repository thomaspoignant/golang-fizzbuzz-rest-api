package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	core "github.com/thomaspoignant/golang-fizzbuzz-rest-api/core"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

// FizzbuzzRequest is the set of parameters you should have to call the api
type FizzbuzzRequest struct {
	String1 string `form:"string1" binding:"required"`
	String2 string `form:"string2" binding:"required"`
	Int1    int    `form:"int1" binding:"required"`
	Int2    int    `form:"int2" binding:"required"`
	Limit   int    `form:"limit" binding:"required"`
}

var initialized = false
var ginLambda *ginadapter.GinLambda

// Handler to wrap gin to lambda
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if !initialized {
		router := SetupRouter()
		ginLambda = ginadapter.New(router)
		initialized = true
	}

	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.Proxy(req)
}

func main() {
	switch runAs := os.Getenv("RUN_AS"); runAs {
	case "lambda":
		log.Info("Run as lambda")
		lambdaRun()
	default:
		log.Info("Run locally")
		localRun()
	}
}

func localRun() {
	router := SetupRouter()
	router.Run() // listen and serve on 0.0.0.0:8080
}

func lambdaRun() {
	lambda.Start(Handler)
}

// SetupRouter determine what to do for each api calls
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// the app listen to /v1/fizzbuzz
	router.GET("/v1/fizzbuzz", func(c *gin.Context) {
		var request FizzbuzzRequest
		if err := c.ShouldBindQuery(&request); err != nil {
			// we cannot bind request return a 400 bad request error
			response := gin.H{"status": http.StatusBadRequest, "error": err.Error()}
			log.WithFields(logrus.Fields{"request": request, "response": response}).Error("Invalid parameters :")
			c.JSON(http.StatusBadRequest, response)
			return
		}

		resp, err := core.FizzBuzz(request.String1, request.String2, request.Int1, request.Int2, request.Limit)
		if err != nil {
			// fizzbuzz return an error something went wrong with parameters return a 400 bad request
			response := gin.H{"status": http.StatusBadRequest, "error": err.Error()}
			log.WithFields(logrus.Fields{"request": request, "response": response}).Error("Invalid parameters :")
			c.JSON(http.StatusBadRequest, response)
			return
		}
		// we received a valid response from fizzbuzz return a http 200
		response := gin.H{"status": http.StatusOK, "value": resp}
		log.WithFields(logrus.Fields{"request": request, "response": response}).Info("Success request :")
		c.JSON(http.StatusOK, response)
	})
	return router
}
