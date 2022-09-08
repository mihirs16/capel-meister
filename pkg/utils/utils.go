package utils

import (
	"capel-meister/pkg/logs"
	"encoding/json"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

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
PathFromRepoName returns the absolute filepath
built from repository's name and GITHUB_PATH.
*/
func PathFromRepoName (name string) string {
    return filepath.Join(
        os.Getenv("GITHUB_PATH"), 
        name,
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
WriteJSONToFile writes the contents of data to a JSON file.
*/
func WriteJSONToFile (file string, data interface{}) {
    encodedData, err := json.MarshalIndent(data, "", "\t");
    logs.Error(err);
    
    err = os.WriteFile(file, encodedData, 0644);
    logs.Error(err);
}

