package main

import (
	"flag"
	"fmt"
	"gRPC/config"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

var configFlag = flag.String("config", "./config.toml", "config path")

// 터미널에 go run . -> main 함수 실행 -> cmd의 app.go(네트워크, 리포지토리,서비스에 대한 객체값 가지고) 실행
func main() {
	flag.Parse()
	fmt.Println("터미널에 go run . config=test", *configFlag)
	config.NewConfig(*configFlag)
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
