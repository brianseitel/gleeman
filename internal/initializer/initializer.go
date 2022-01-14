package initializer

import (
	"embed"
	_ "embed"
)

//go:embed public/assets/*
//go:embed tales/*.yaml
//go:embed tales/layout/*
var fs embed.FS

func Fetch(path string) []byte {
	data, err := fs.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return data
}
