package main

import (
	"github.com/kourtzaridisr88/gator/internal/config"
	"github.com/kourtzaridisr88/gator/internal/database"
)

type State struct {
	Config *config.Config
	DB     *database.Queries
}
