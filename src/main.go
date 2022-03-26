package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting up server...")
	gin.ForceConsoleColor()
	r := gin.Default()
	r.Run()
}
