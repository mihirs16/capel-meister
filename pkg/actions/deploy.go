package actions

import (
	"capel-meister/pkg/container"
	"capel-meister/pkg/repository"
)


/*
DeployAll deploys all the repositories specified to the server.
*/
func Deploy (deployMetadata []repository.RepoMetadata) {
    
    // clone repositories and checkout branches
    for _, repoMetadata := range deployMetadata {
        repoMetadata.LoadRepository();
        repoMetadata.CheckoutBranch();
        repoMetadata.PullRepository();
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
