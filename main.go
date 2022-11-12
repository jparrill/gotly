package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/jparrill/gotly/internal/config"
	"github.com/jparrill/gotly/internal/server"
	"github.com/jparrill/gotly/internal/utils"
)

func main() {
	// Set Context
	ctx := context.Background()

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

	// Recover the configuration
	config.RecoverConfig(dir + "/assets/samples/config.yaml")
	// Initialize the logger
	config.InitLogger()
	config.MainLogger.Info("Initializing DDBB")

	// Run Server
	server.Run(ctx, dir, &config.CFG)
}
