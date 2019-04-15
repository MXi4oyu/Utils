package main

import (
	"context"
	"fmt"
	"github.com/MXi4oyu/Utils/cnencoder/gbk"
	"github.com/MXi4oyu/Utils/gziputil"
	"github.com/MXi4oyu/Utils/subprocess"
	"github.com/MXi4oyu/Utils/walkpath"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	str, err := subprocess.RunCommand(ctx, "sh", "-c", "ls")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(gbk.Decode(str))

	//获取指定目录下的所有文件，不进入子目录，可通过后缀名过滤
	files, err := walkpath.ListDir("/var/www/", ".php")
	fmt.Println(files, err)

	//获取指定目录及其子目录下的所有文件，可通过后缀名进行过滤
	filesall, err := walkpath.WalkDir("/var/www", ".php")
	fmt.Println(filesall, err)

	//解压tar文件

	gziputil.TarDeCompress("/tmp/layer.tar", "/tmp/testtar/")

}
