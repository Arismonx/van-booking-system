package main

import (
	"fmt"

	"github.com/Arismonx/van-booking-system/config"
)

func main() {
	conf := config.LoadConfig() // test config
	fmt.Println(conf)
}
