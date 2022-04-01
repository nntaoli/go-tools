package file

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
	"path"
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
