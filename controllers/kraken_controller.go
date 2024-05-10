package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func KrakenGetInfoHealth(ctx *gin.Context) {
	// Status URL of Kraken Server
	url := "https://api.kraken.com/0/public/SystemStatus"

	// Make the call to the endpoint
	respuesta, err := http.Get(url)
	if err != nil {
		fmt.Println("Error calling the API:", err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	defer respuesta.Body.Close()

	// Read the API response
	cuerpoRespuesta, err := io.ReadAll(respuesta.Body)
	if err != nil {
		fmt.Println("Error reading the response:", err)
		return
	}

	// Print the response
	fmt.Println("Respuesta de la API:", string(cuerpoRespuesta))

	// Transform the API response into JSON
	var jsonData map[string]interface{}
	err = json.Unmarshal(cuerpoRespuesta, &jsonData)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return JSON response
	ctx.JSON(http.StatusOK, jsonData)
}


