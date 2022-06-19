package main

import (
	"embed"

	"github.com/StarlaneStudios/EtherealGit/backend"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	backend.Initialize(assets)
}
