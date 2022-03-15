package subpackage

import (
	"rootpkg"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	configFile, err := rootpkg.RootFS.ReadFile("config/env")
	if err != nil {
		t.Fatalf("unable to open config/env file: %s", err)
	}

	if string(configFile) != "FOO=bar\n" {
		t.Fatalf("config file contents differ from expected: %s", string(configFile))
	}
}
