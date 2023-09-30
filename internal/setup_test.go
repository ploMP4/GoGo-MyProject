package internal

import (
	"os"
	"syscall"
	"testing"

	"github.com/fatih/color"
)

func TestMain(m *testing.M) {
	color.Output = os.NewFile(uintptr(syscall.Stdin), os.DevNull)

	os.Exit(m.Run())
}
