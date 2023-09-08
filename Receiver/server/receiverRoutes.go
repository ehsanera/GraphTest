package server

import (
	"Receiver/customMiddleware"
	"Receiver/models"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

func receiverRoutes(e *echo.Echo) {
	e.Use(middleware.BodyLimit("8K"))
	e.Use(customMiddleware.MinBodySizeMiddleware(50))

	newLimiter := tollbooth.NewLimiter(10000, &limiter.ExpirableOptions{
		DefaultExpirationTTL: time.Second,
	})
	newLimiter.SetMessage("Too Many Requests")

	e.POST("/send", postHandler, customMiddleware.RateLimitMiddleware(newLimiter))
}

func postHandler(c echo.Context) error {
	requestBody, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		return err
	}

	requestBodyString := string(requestBody)

	request := &models.SendRequest{
		Message: requestBodyString,
	}

	fmt.Printf("Receive: %d\n", len(request.Message))

	message := models.Message{
		Message:  request.Message,
		Received: false,
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}

	err = sendDataToSocket(jsonData)

	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

func sendDataToSocket(data []byte) error {
	socketServerAddr := "localhost:8081"

	conn, err := net.Dial("tcp", socketServerAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	_, err = writer.Write(data)
	if err != nil {
		return err
	}

	return nil
}
