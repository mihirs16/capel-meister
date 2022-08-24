package main

import (
	"os"
	"capel-meister/internal/server"
	"capel-meister/pkg/logs"
	"github.com/joho/godotenv"
);


func main () {
    // load dotenv
    err := godotenv.Load();
    logs.Error(err);

    // server
    server.BuildRouter(os.Getenv("GIN_MODE"));
}
