/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import "github.com/leijeng/huo-admin/cmd"

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.io,direct
//go:generate go mod tidy
//go:generate go mod download
//go:generate swag init --parseDependency --parseDepth=6

// @title Huo API
// @version V1.0.1
// @description 致力于做一个开发快速，运行稳定的框架
// @contact.name   lei
// @contact.url    https://github.com/leijeng/huo-admin
// @contact.email  lisite199505@gmail.com

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @BasePath /api/v1
func main() {
	cmd.Execute()
}
