package main

import "github.com/cf-furnace/k8s-cc-uploader/cmd"

var version = ""

func main() {
	cmd.Execute(version)
}
