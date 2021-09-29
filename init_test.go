package main

import (
	"os"
	"testing"

	"github.com/yaoapp/gou"
	"github.com/yaoapp/xiang/config"
)

var cfg config.Config

func TestMain(m *testing.M) {

	// Run test suites
	exitVal := m.Run()

	// we can do clean up code here
	gou.KillPlugins()

	os.Exit(exitVal)
}
