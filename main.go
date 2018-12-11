// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"gopkg.in/src-d/go-git.v4"
	"fmt"
	ssh2"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"os"
	"os/user"
)

func main() {
	u, e := user.Current()

	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(u.HomeDir)
	sshAuth, err := ssh2.NewPublicKeysFromFile("git", u.HomeDir+"/.ssh/id_rsa", "password")

	if err != nil {
		fmt.Println(err)
	}

	git.PlainClone("~/.basket", false, &git.CloneOptions{
		URL: "git@github.com/basketsavings/massclarity-platform.git",
		Auth: sshAuth,
		Progress:os.Stdout,
	})
}
