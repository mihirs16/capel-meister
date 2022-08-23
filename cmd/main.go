package main

import (
	"fmt"
	"os"
	// "capel-meister/internal/server"
	"capel-meister/pkg/logs"
	"capel-meister/pkg/utils"
	"github.com/joho/godotenv"
    "capel-meister/pkg/container"
);


func main () {
    // load dotenv
    err := godotenv.Load();
    logs.Error(err);

    // server
    // server.BuildRouter(os.Getenv("GIN_MODE"));

    composeMetadata := container.ComposeMetadata{};
    utils.LoadJSONFromFile(os.Getenv("DOCKER_COMPOSE_FILE_PATH"), &composeMetadata)
    fmt.Println(composeMetadata)
}
