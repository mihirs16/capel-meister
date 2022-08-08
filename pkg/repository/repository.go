package repository

import (
	"capel-meister/pkg/logs"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

/*
RepoMetadata has metadata
for the repository, worktree and submodules
*/
type RepoMetadata struct {
    repository  *git.Repository
    worktree    *git.Worktree
    submodules  git.Submodules
};

func LoadRepository (repoPath string) *RepoMetadata {
    fmt.Println(repoPath);
    repo, err := git.PlainOpen(repoPath);
    logs.Error(err);

    worktree, err := repo.Worktree();
    logs.Error(err);

    submodules, err := worktree.Submodules();
    logs.Error(err);

    repoMetadata := RepoMetadata{repo, worktree, submodules};
    return &repoMetadata;
}

func (repoMetadata *RepoMetadata) PullRepository () {
    err := repoMetadata.worktree.Checkout(&git.CheckoutOptions{
        Branch: "refs/heads/main",
    });
    logs.Error(err);
    err = repoMetadata.worktree.Pull(&git.PullOptions{
        RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
    });
    if err == git.NoErrAlreadyUpToDate {
        logs.Info("Already up to date.");
    } else {
        logs.Error(err);
    }
}

func (repoMetadata *RepoMetadata) UpdateSubmodules () { 
    err := repoMetadata.submodules.Update(&git.SubmoduleUpdateOptions{
        Init: true,
        RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
    });
    logs.Error(err);

    for i := range repoMetadata.submodules {
        tempRepoMetadata := LoadRepository(os.Getenv("MEISTER_PATH") + repoMetadata.submodules[i].Config().Path);
        tempRepoMetadata.PullRepository();
    }

}
