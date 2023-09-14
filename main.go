package main

import (
	"gin/cmd/server"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Printf("[musiku-main] failed to load .env file. Error : %v\n", err)
	// 	return
	// }

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		os.Exit(server.Run())
	}()
	wg.Wait()
}
