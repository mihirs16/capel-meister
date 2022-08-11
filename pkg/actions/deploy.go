package actions

import (
	"capel-meister/pkg/repository"
)

/*
Deploy deploys all the repositories specified to the server
*/
func Deploy (DeployMetadata map[string]string) {
    
    // clone repositories and checkout branches
    for repoUrl, branchName := range DeployMetadata {
        repoMetadata := repository.CloneRepository(repoUrl);
        repoMetadata.CheckoutBranch(branchName);
    };

    // pull capel-meister
    meisterRepo := repository.LoadRepository("https://github.com/mihirs16/capel-meister");
    meisterRepo.PullRepository();

    // docker compose -f choir.yml up --build --detach --remove-orphans
}
