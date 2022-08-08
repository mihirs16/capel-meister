package main

import (
    "os"
	"capel-meister/pkg/logs"
	"capel-meister/pkg/repository"
	"github.com/joho/godotenv"
);

func main () {
    // load dotenv
    err := godotenv.Load();
    logs.Error(err);
    
    // repository operations
    repoMetadata := repository.LoadRepository(os.Getenv("MEISTER_PATH"));
    repoMetadata.PullRepository();
    repoMetadata.UpdateSubmodules();
}
