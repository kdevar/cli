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

type CliConfig struct {
	session                     *session.Session
	Cloudformation              *cloudformation.CloudFormation
	InfrastructurePath          string
	PlatformTemplatePath        string
	InfrastructureRepositoryUrl string
	PlatformTemplateUrl         string
	PemPath                     string
	Auth                        ssh.AuthMethod
}

func NewCliConfig() *CliConfig {
	u, e := user.Current()

	if e != nil {
		fmt.Println(e)
	}

	pem := u.HomeDir + pemPath

	sshAuth, err := ssh.NewPublicKeysFromFile("git", pem, "password")
	if err != nil {
		fmt.Println(err)
	}
	s, _ := session.NewSession(&aws.Config{})
	return &CliConfig{
		session:                     s,
		Cloudformation:              cloudformation.New(s),
		InfrastructureRepositoryUrl: "git@github.com/basketsavings/massclarity-platform.git",
		Auth:                        sshAuth,
	}
}
