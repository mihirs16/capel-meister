package main

import (
	// "os"
    "github.com/joho/godotenv"
	"capel-meister/pkg/logs"
	// "capel-meister/pkg/repository"
);

func main () {
    // load dotenv
    err := godotenv.Load();
    logs.Error(err);
    
}
