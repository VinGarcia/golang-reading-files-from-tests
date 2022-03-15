package rootpkg

import (
	"embed"
)

//go:embed config/*
var root embed.FS

func ReadFile(filepath string) ([]byte, error) {
	return root.ReadFile(filepath)
}
