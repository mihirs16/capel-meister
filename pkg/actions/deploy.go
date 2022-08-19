package actions

import (
	"capel-meister/pkg/container"
	"capel-meister/pkg/logs"
	"capel-meister/pkg/repository"
)

/*
Deploy deploys all the repositories specified to the server
*/
func Deploy (file string) {
    deployMetadata := repository.LoadMetadata(file);
    
    // clone repositories and checkout branches
    for _, repoMetadata := range deployMetadata {
        repoMetadata.LoadRepository();
        repoMetadata.CheckoutBranch();
    };

    // pull capel-meister
    meisterRepo := repository.RepoMetadata{
        URL: "https://github.com/mihirs16/capel-meister",
        BranchName: "main",
    };
    meisterRepo.LoadRepository();
    meisterRepo.PullRepository();

    // docker compose 
    container.Build();
    container.Up();
}
