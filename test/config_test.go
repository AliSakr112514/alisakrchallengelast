package test

import (
	"Challenge/config"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config := config.LoadConfig()
	if config == nil {
		t.Errorf("The test failed during getting configurations")
	}

}
