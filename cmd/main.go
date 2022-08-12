package main

import (
	// "os"
    "github.com/joho/godotenv"
	"capel-meister/pkg/logs"
    "capel-meister/pkg/actions"
);

func main () {
    // load dotenv
    err := godotenv.Load();
    logs.Error(err);
    
    // deploy action
    actions.Deploy(
        map[string]string{
            "https://github.com/mihirs16/monte-carlo": "main",
        },
    );
}
