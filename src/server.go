package main

import (
	"log"
	"net/http"
	"os"

	"github.com/BlankDev117/gomockserver/src/routes"
)

// #region Variables

const baseEndpoint = "/"

const portParamName = "PORT"
const filePathParamName = "CONFIG_PATH"
const defaultDataPath = "/app/data"

// #endregion

func main() {
	startWebServer()
}

/* #region Helpers */

func startWebServer() {

	port, filePath := getServerConfiguration()

	log.Printf("Checking for api configuration at %s", filePath)

	routeMap, err := ReadConfigurationFiles(filePath)

	if err != nil {
		panic(err)
	}

	manager := routes.NewRouteManager(routeMap)

	http.HandleFunc(baseEndpoint, manager.ProcessAPIRequest)

	log.Printf("Starting gomockserver. Listening on endpoints [%s] and port %s", baseEndpoint, port)

	err = http.ListenAndServe(":"+port, nil)

	if err != nil {
		panic(err)
	}
}

func getServerConfiguration() (string, string) {

	port := os.Getenv(portParamName)

	if len(port) == 0 {
		panic("Environment Port not set")
	}

	filePath := os.Getenv(filePathParamName)

	if len(filePath) == 0 {
		filePath = defaultDataPath
	}

	return port, filePath
}

/* #endregion */
