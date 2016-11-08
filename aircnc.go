package main

import (
	"github.com/matheusvill/aircnc/api"
	"github.com/matheusvill/aircnc/storage"
)

func main() {
	storage.CreateDb()
	api.Start()
}
