package server

import (
	"Receiver/customCache"
	"Receiver/customMiddleware"
	"Receiver/models"
	"Receiver/socket"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/ioutil"
	"net/http"
)

func receiverRoutes(e *echo.Echo) {
	e.Use(middleware.BodyLimit("8K"))
	e.Use(customMiddleware.MinBodySizeMiddleware(50))

	e.POST("/send", postHandler)
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

	fmt.Println("Receive: " + request.Message)

	message := customCache.Message{
		Message:  request.Message,
		Received: false,
	}
	err = message.Create(context.Background(), customCache.Db, "messages", message)
	if err != nil {
		panic(err)
	}

	socket.Check()

	return c.JSON(http.StatusOK, nil)
}
