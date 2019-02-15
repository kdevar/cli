package config

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"os/user"
)

var pemPath = "/.ssh/id_rsa"

type Env string

func (e *Env) String() string{
	return string(*e)
}

var (
	Sandbox          Env = "Sandbox"
	Dev              Env = "Dev"
	Stage            Env = "Stage"
	Prod             Env = "Prod"
	MassclarityDev   Env = "MassclarityDev"
	MassclarityStage Env = "MassclarityStage"
	MassclarityProd  Env = "MassclarityProd"
)


type CliConfig struct {
	session                     *session.Session
	Cloudformation              *cloudformation.CloudFormation
	InfrastructurePath          string
	PlatformTemplatePath        string
	InfrastructureRepositoryUrl string
	PlatformTemplateUrl         string
	PemPath                     string
	Auth                        ssh.AuthMethod
	Environment					Env
}

func NewCliConfig() *CliConfig {
	region := "us-east-1"
	u, e := user.Current()
	if e != nil {
		fmt.Println(e)
	}

	pem := u.HomeDir + pemPath

	sshAuth, err := ssh.NewPublicKeysFromFile("git", pem, "password")
	if err != nil {
		fmt.Println(err)
	}
	s, _ := session.NewSession(&aws.Config{
		Region: &region,
	})
	return &CliConfig{
		session:                     s,
		Cloudformation:              cloudformation.New(s),
		InfrastructurePath:          u.HomeDir + "/.basket/massclarity-infrastructure",
		InfrastructureRepositoryUrl: "git@github.com:basketsavings/massclarity-infrastructure.git",
		Auth:                        sshAuth,
	}
}
