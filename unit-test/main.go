package main

import "github.com/afrianska/playgo/unit-test/pkg/utils"

func main() {
	app := utils.Setup()

	app.Listen(":3000")
}