package walkpath

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//获取指定目录下的所有文件，不进入子目录，可通过后缀名过滤
func ListDir(dirpath, suffix string) (files []string, err error) {

	files = make([]string, 0, 30)
	dir, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return nil, err
	}

	pathsep := string(os.PathSeparator)

	suffix = strings.ToUpper(suffix)

	for _, fi := range dir {

		if fi.IsDir() {
			continue
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, dirpath+pathsep+fi.Name())
		}
	}

	return files, err

}

//获取指定目录及其子目录下的所有文件，可通过后缀名进行过滤
func WalkDir(dirpath, suffix string) (files []string, err error) {

	files = make([]string, 0, 30)

	suffix = strings.ToUpper(suffix) //不区分文件后缀名大小写

	err = filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() { //忽略目录
			return nil
		}

		if strings.HasSuffix(strings.ToUpper(info.Name()), suffix) {
			files = append(files, path)
		}

		return nil

	})

	return files, err

}
