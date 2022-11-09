package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jparrill/gotly/internal/server"
	"github.com/jparrill/gotly/internal/utils"
)

func main() {
	// Locate basepath for repository
	dir, err := utils.GetBaseRepoPath(filepath.Dir(os.Args[0]))
	if err == os.ErrNotExist {
		local, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		dir, err = utils.GetBaseRepoPath(local)
		if err != nil {
			log.Fatal("Basepath for repo not found, please make sure you are located in the right place")
		}
	}

	// TODO: (cobra and maybe viper) Load config from file
	fmt.Println(dir)

	// Run Server
	server.Run(dir)
}
