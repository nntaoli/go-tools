package file

import (
	"os"
	"testing"
)

func TestPathExists(t *testing.T) {
	t.Log(IsPathExists("File.go"))
}

func TestGetFileSizeKB(t *testing.T) {
	t.Log(GetFileSizeKB("File.go"))
}

func TestGetFileMd5(t *testing.T) {
	f, _ := os.Open("File.go")
	t.Log(GetFileMd5(f))
}

func TestEnsureCreateFile(t *testing.T) {
	t.Log(EnsureCreateFile("a/b/c/xx.go"))
}

func TestGetProgramRunningDir(t *testing.T) {
	t.Log(GetProgramRunningDir())
	t.Log(os.UserHomeDir())
}
