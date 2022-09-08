package repository

import (
	"capel-meister/pkg/logs"
	"github.com/go-git/go-git/v5"
    "capel-meister/pkg/utils"
)


/*
RepoMetadata has metadata
for the repository, worktree and submodules
*/
type RepoMetadata struct {
    repository  *git.Repository
    // worktree    *git.Worktree
    URL         string          `json:"url"`
    BranchName  string          `json:"branch"`
    Services    map[string]struct {
        Build   string      `json:"build"`
        Ports   []string    `json:"ports"`
    }                       `json:"services"` 
};


/*
InitRepository initializes an empty repo
*/
func InitRepository (name string) *RepoMetadata {
    repoPath := utils.PathFromRepoName(name)
    repo, err := git.PlainInit(repoPath, true);
    logs.Error(err);

    return &RepoMetadata{ repository: repo };
}


/* 
CloneRepository clones the repo 
from the `repoUrl` specified 
to the repoPath specified
func (repoMetadata *RepoMetadata) CloneRepository () {
    repoPath := utils.PathFromRepoURL(repoMetadata.URL);
    repo, err := git.PlainClone(repoPath, true, &git.CloneOptions{ 
        URL: repoMetadata.URL,
        RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
    });
    logs.Error(err);

    repoMetadata.repository = repo;
}
*/


/*
LoadRepository loads the repo if `.git` dir
if found in the `repoPath`

func (repoMetadata *RepoMetadata) LoadRepository () {
    repoPath := utils.PathFromRepoURL(repoMetadata.URL);
    repo, err := git.PlainOpen(repoPath);
    if err == git.ErrRepositoryNotExists {
        repoMetadata.CloneRepository();
        return;
    }
    logs.Error(err);

    worktree, err := repo.Worktree();
    logs.Error(err);

    repoMetadata.repository = repo;
    repoMetadata.worktree   = worktree;
}
*/

/* 
PullRepository pulls the repository represented by RepoMetadata
func (repoMetadata *RepoMetadata) PullRepository () {
    err := repoMetadata.worktree.Pull(&git.PullOptions{
        RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
    });
    if err == git.NoErrAlreadyUpToDate {
        logs.Info("Already up to date.");
    } else {
        logs.Error(err);
    }
}
*/

/*
CheckoutBranch checks out the branch for the repo
func (repoMetadata *RepoMetadata) CheckoutBranch () {
    branch, err := repoMetadata.repository.Branch(repoMetadata.BranchName);
    logs.Error(err);

    err = repoMetadata.worktree.Checkout(&git.CheckoutOptions{
        Branch: branch.Merge,
    });
    logs.Error(err);
}
*/

