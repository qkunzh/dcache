package main

import "github.com/qkunzh/dcache"

func main() {
	cacheServer := dcache.NewCacheServer("localhost", 9999, 1024*1024)
	cacheServer.Start()
}
