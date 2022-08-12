package container

import (
	"capel-meister/pkg/logs"
	"os/exec"
);

/*
Build builds all the services' images in docker-compose.yml
*/
func Build () {
    buildCmd := exec.Command("docker", "compose", "build");
    stdOut, err := buildCmd.CombinedOutput();
    logs.Info(string(stdOut));
    logs.Error(err);
}

/*
Up updates and runs the built images of services in the docker-compose.yml
*/
func Up () {
    upCmd := exec.Command("docker", "compose", "up", "--detach"); 
    stdOut, err := upCmd.CombinedOutput();
    logs.Info(string(stdOut));
    logs.Error(err);
}
