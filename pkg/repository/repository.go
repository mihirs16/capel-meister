package repository

import (
	"capel-meister/pkg/logs"
	"github.com/go-git/go-git/v5"
)


/*
RepoMetadata has metadata
for the repository, worktree and submodules
*/
type RepoMetadata struct {
    repository  *git.Repository
    worktree    *git.Worktree
    URL         string          `json:"url"`
    BranchName  string          `json:"branch"`
};


/* 
CloneRepository clones the repo 
from the `repoUrl` specified 
to the repoPath specified
*/
func (repoMetadata *RepoMetadata) CloneRepository () {
    repoPath := PathFromRepoURL(repoMetadata.URL);
    repo, err := git.PlainClone(repoPath, false, &git.CloneOptions{ 
        URL: repoMetadata.URL,
        RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
    });
    logs.Error(err);

    worktree, err := repo.Worktree();
    logs.Error(err);

    repoMetadata.repository = repo;
    repoMetadata.worktree   = worktree;
}


/*
LoadRepository loads the repo if `.git` dir
if found in the `repoPath`
*/
func (repoMetadata *RepoMetadata) LoadRepository () {
    repoPath := PathFromRepoURL(repoMetadata.URL);
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


/* 
PullRepository pulls the repository represented by RepoMetadata
*/
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


/*
CheckoutBranch checks out the branch for the repo
*/
func (repoMetadata *RepoMetadata) CheckoutBranch () {
    branch, err := repoMetadata.repository.Branch(repoMetadata.BranchName);
    logs.Error(err);

    err = repoMetadata.worktree.Checkout(&git.CheckoutOptions{
        Branch: branch.Merge,
    });
    logs.Error(err);
}

