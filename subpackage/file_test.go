package subpackage

import (
	"rootpkg"
	"testing"
)

func TestOpenConfig(t *testing.T) {
	configFile, err := rootpkg.RootFS.ReadFile("config/env")
	if err != nil {
		t.Fatalf("unable to open config/env file: %s", err)
	}

	if string(configFile) != "FOO=bar\n" {
		t.Fatalf("config file contents differ from expected: %s", string(configFile))
	}
}

func TestOpenHTML(t *testing.T) {
	configFile, err := rootpkg.RootFS.ReadFile("html/index.html")
	if err != nil {
		t.Fatalf("unable to open html/index.html file: %s", err)
	}

	if string(configFile) != "<body>example</body>\n" {
		t.Fatalf("html file contents differ from expected: %s", string(configFile))
	}
}
