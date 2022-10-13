package file

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

func IsPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetFileSizeKB(filePath string) (size float64, err error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	size = float64(info.Size()) / 1000.0
	return
}

func GetFileMd5(src io.Reader) string {
	hash := md5.New()
	io.Copy(hash, src)
	return hex.EncodeToString(hash.Sum(nil))
}

// EnsureCreateFile
//所有目录会一起创建
func EnsureCreateFile(f string) error {
	dirPath := path.Dir(f)
	log.Println(dirPath)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}
	_, err = os.Create(f)
	return err
}

//GetProgramRunningDir
//获取程序运行目录
func GetProgramRunningDir() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return dir, nil
}

func IsEmptyDir(dir string) (bool, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return false, err
	}
	if len(files) == 0 {
		return true, nil
	}
	return false, nil
}
