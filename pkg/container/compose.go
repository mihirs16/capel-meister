package container

import (
	"capel-meister/pkg/logs"
	"os/exec"
);


/*
ComposeMetadata has metadata for the 
docker compose YAML file
*/
type ComposeMetadata struct {
    Version     string      `json:"version"`
    Services    map[string]struct {
        Build   string      `json:"build"`
        Ports   []string    `json:"ports"`
    }                       `json:"services"` 
};


/*
Build builds all the services' images in docker-compose.yml
*/
func Build () {
    executeCommand("docker", "compose", "build");
}


/*
Up updates and runs the built images of services in the docker-compose.yml
*/
func Up () {
    executeCommand("docker", "compose", "up", "--detach"); 
}


/*
Down stops all running services in docker-compose.yml
*/
func Down (service string) {
    executeCommand("docker", "compose", "down");
}


/*
executeCommand is a helper function for executing commands passed as arguments.
*/
func executeCommand (cmd string, args ...string) {
    toExecCmd := exec.Command(cmd, args...);
    stdOut, err := toExecCmd.CombinedOutput();
    logs.Info(string(stdOut));
    logs.Error(err);
}

