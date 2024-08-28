package main

import (
	"backend-server/init/cmd"
	"flag"
	"fmt"
	"net/http"
)

var configPathFlag = flag.String("config", "./config.toml", "config file not found")

func main() {
	// http.HandleFunc("/", helloWorld)

	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	fmt.Println("에러 발생!")
	// 	panic(err)
	// }
	flag.Parse()

	cmd.NewCmd(*configPathFlag)
	fmt.Println("서버종료")
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello GoWeb!")
}
