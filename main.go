package main

import (
	"./services"
	"fmt"
)

func main() {
	res := services.GetKLine("btcusdt", "1min", 200)
	fmt.Println(res)
}
