package main

import (
	"eye/internal/initialize"
)

func main(){
	initialize.NewEye("./configs/application.toml").InitAll()
}
