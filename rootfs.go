package rootpkg

import (
	"embed"
)

//go:embed config/*
var RootFS embed.FS
