package repository

import (
    "os"
    "errors"
    "strings"
    "net/url"
    "path/filepath"
	"encoding/json"
    "capel-meister/pkg/logs"
);


/*
PathFromRepoURL returns the absolute filepath
built from repository's URL and GITHUB_PATH.
*/
func PathFromRepoURL (repoUrl string) string {
    parsedUrl, err := url.Parse(repoUrl);
    logs.Error(err);

    return filepath.Join(
        os.Getenv("GITHUB_PATH"), 
        strings.Split(parsedUrl.Path, "/")[2],
    );
}


/*
LoadJSONFromFile loads the contents of a JSON file into data.
*/
func LoadJSONFromFile (file string, data interface{}) {
    encodedData, err := os.ReadFile(file);
    logs.Error(err);
    
    err = json.Unmarshal([]byte(encodedData), data);
    logs.Error(err);
}


/*
WriteJSONToFile writes the contents to a JSON file
*/
func WriteJSONToFile (file string, data interface{}) {
    encodedData, err := json.MarshalIndent(data, "", "\t");
    logs.Error(err);
    
    err = os.WriteFile(file, encodedData, 0644);
    logs.Error(err);
}

