package main

import "github.com/CyberMidori/gin-ranking/router"

func main() {
	r := router.Router()

	r.Run(":9999")
}
