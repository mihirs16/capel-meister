package main

import (
	"capel-meister/pkg/logs"
	"capel-meister/pkg/repository"
	"os"
	"github.com/joho/godotenv"
);

func main () {
    // load dotenv
    err := godotenv.Load();
    logs.Error(err);
    
    // open this repository
    repository.LoadRepository(os.Getenv("MEISTER_PATH"));
}
