package rootpkg

import (
	"embed"
)

// You can add more subdirectories here, by separating them with spaces
// like we do for the config and html directories below:
//
//go:embed config/* html/*
var RootFS embed.FS
