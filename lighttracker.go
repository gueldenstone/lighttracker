package main

import (
	"context"
	"strconv"
)

func init() {}

var ctx = context.Background()

const port = 8080

func main() {

	serveFrontend("frontend/dist")

	connectRedis("localhost:6372", "")

	router.Run(":" + strconv.Itoa(port))
}
