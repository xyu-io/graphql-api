package main

import (
	"fmt"
	"graphql-api/internal/servers"
	"os"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	ginServ := servers.GetInstance()
	err := ginServ.Start()
	if err != nil {
		fmt.Println("项目启动失败.")
	}

	<-signalChan
}
