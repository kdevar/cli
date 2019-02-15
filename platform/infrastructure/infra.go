package infrastructure

import (
	"fmt"
	"github.com/kdevar/cli/config"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"log"
)

type InfrastructureManager struct {
	Config     *config.CliConfig
	Repository *git.Repository
}

func (r *InfrastructureManager) Init() *InfrastructureManager {

	err := r.openRepo()

	if err == git.ErrRepositoryNotExists {
		e := r.cloneRepo()
		CheckIfError(e)
	}

	return r
}

func (r *InfrastructureManager) cloneRepo() error {
	Info("cloning repo in to %v from %v ", r.Config.InfrastructurePath, r.Config.InfrastructureRepositoryUrl)
	repo, e := git.PlainClone(r.Config.InfrastructurePath, false, &git.CloneOptions{
		URL:      r.Config.InfrastructureRepositoryUrl,
		Auth:     r.Config.Auth,
		Progress: os.Stdout,
	})

	if e != nil {
		Info("error %v", e)
		return e
	}
	r.Repository = repo
	return nil
}

func (r *InfrastructureManager) openRepo()  error {
	Info("opening repo %v", r.Config.InfrastructurePath)
	repo, e := git.PlainOpen(r.Config.InfrastructurePath)
	if e != nil {
		Info("error %v", e)
		return e
	}
	r.Repository = repo
	return nil
}

func (r *InfrastructureManager) GetLatest() *InfrastructureManager {
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

func (r *InfrastructureManager) PublishParams() *InfrastructureManager {
	path := r.Config.InfrastructurePath + "stacks/params/cloudformation.yaml"
	stack := config.Sandbox.String() + "-params"

	_, err := r.Config.Cloudformation.CreateStack(&cloudformation.CreateStackInput{
		TemplateURL: &path,
		StackName: &stack,
	})

	if err != nil {
		log.Fatalf("error %v", err)
	}

	return r
}

func (r *InfrastructureManager) PublishBase() *InfrastructureManager {
	Info("publishing base")
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


func NewInfrastructureManager(config *config.CliConfig) *InfrastructureManager {
	return &InfrastructureManager{
		Config: config,
	}
}