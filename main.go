package main

import (
	"fmt"
	"context"
	"time"
	"github.com/MXi4oyu/Utils/subprocess"
	"github.com/MXi4oyu/Utils/cnencoder/gbk"
)

func main()  {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(60)*time.Second)
	defer cancel()

	str,err:=subprocess.RunCommand(ctx,"sh","-c","ipconfig")
	if err!=nil{
		fmt.Println(err.Error())
	}

	fmt.Println(gbk.Decode(str))

}
