package gziputil

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"strings"
)

func unTarFile(dstFile string, tr *tar.Reader) error {
	// 创建空文件，准备写入解包后的数据
	fw, er := os.Create(dstFile)
	if er != nil {
		return er
	}
	defer fw.Close()

	// 写入解包后的数据
	_, er = io.Copy(fw, tr)
	if er != nil {
		return er
	}

	return nil
}

//解压tar文件
func TarDeCompress(tarFile, dstFile string) error {

	srcFile, err := os.Open(tarFile)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	defer srcFile.Close()

	tr := tar.NewReader(srcFile)

	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err.Error())
				return err
			}
		}

		fname := hdr.Name
		finfo := hdr.FileInfo()

		dstfullpath := dstFile + fname

		if hdr.Typeflag == tar.TypeDir {
			//创建目录
			os.MkdirAll(dstfullpath, finfo.Mode().Perm())
			//设置目录权限
			os.Chmod(dstfullpath, finfo.Mode().Perm())
		} else {
			//创建文件所在目录
			fpath := dstfullpath[:strings.LastIndex(dstfullpath, "/")+1]
			os.MkdirAll(fpath, os.ModePerm)
			//将tr中的数据写入到文件中

			if err := unTarFile(dstfullpath, tr); err != nil {
				//log.Fatal(err.Error())
				return err
			}

			//设置文件权限
			os.Chmod(dstfullpath, finfo.Mode().Perm())

		}

	}

	return nil
}
