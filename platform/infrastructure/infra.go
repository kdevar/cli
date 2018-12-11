package infrastructure

import (
	"fmt"
	"github.com/kdevar/cli/config"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

type InfrastructureRepository struct {
	Config     *config.CliConfig
	Path       string
	Url        string
	Repository *git.Repository
}

func (r *InfrastructureRepository) Init() *InfrastructureRepository {
	err := r.open()

	if err != nil {
		err = r.clone()
	}
	return r
}

func (r *InfrastructureRepository) clone() error {
	repo, e := git.PlainClone(r.Config.InfrastructurePath, false, &git.CloneOptions{
		URL:      r.Config.InfrastructureRepositoryUrl,
		Auth:     r.Config.Auth,
		Progress: os.Stdout,
	})

	if e != nil {
		return e
	}
	r.Repository = repo
	return nil
}

func (r *InfrastructureRepository) open()  error {
	repo, e := git.PlainOpen(r.Config.InfrastructurePath)
	if e != nil {
		return e
	}
	r.Repository = repo
	return nil
}

func (r *InfrastructureRepository) GetLatest() *InfrastructureRepository {
	po := &git.PullOptions{
		Auth:       r.Config.Auth,
		RemoteName: "Origin",
	}

	w, err := r.Repository.Worktree()

	if err != nil {

	}

	w.Pull(po)
	return r
}

func (r *InfrastructureRepository) PublishParams() *InfrastructureRepository {
	path := r.Config.InfrastructurePath + "stacks/base/cloudformation.yaml"

	_, err := r.Config.Cloudformation.CreateStack(&cloudformation.CreateStackInput{
		TemplateURL: &path,
	})

	if err != nil {

	}

	return r
}

func (r *InfrastructureRepository) PublishBase() *InfrastructureRepository {
	return r
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
