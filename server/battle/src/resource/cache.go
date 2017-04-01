package resource

import (
	"io/ioutil"
	"os"
	"path"
)

func writeCacheFile(fileName string, data []byte) error {
	filePath := path.Join(cacheDir, fileName)
	if err := os.MkdirAll(path.Dir(filePath), os.ModePerm); err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, os.ModePerm)
}

func existsCacheFile(fileName string) (string, bool) {
	filePath := path.Join(cacheDir, fileName)
	_, err := os.Stat(filePath)
	return filePath, err == nil
}
