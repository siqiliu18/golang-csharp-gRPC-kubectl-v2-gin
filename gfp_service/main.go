package main

import (
	"fmt"

	"gfpservice/gfpfuncs"
)

func main() {
	fmt.Println("[main] gfp service")
	svc, err := gfpfuncs.NewService()
	if err != nil {
		fmt.Println("[main] failed to create service;", svc.Name, err.Error())
	}
}