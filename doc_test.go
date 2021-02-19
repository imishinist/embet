package embet_test

import (
	"embed"
	"log"
	"os"

	"github.com/imishinist/embet"
)

//go:embed assets/*
var assets embed.FS

func ExampleList() {
	list, err := embet.List(asset, "assets")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range list {
		fmt.Println(f)
	}
}

func ExampleWriteEmbedFiles() {
	dest := "dest"
	if err := os.Mkdir(dest, 0755); err != nil {
		log.Fatal(err)
	}
	if err := embet.WriteEmbedFiles(assets, "assets", dest); err != nil {
		log.Fatal(err)
	}
}
