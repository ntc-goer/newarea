package main

import "newarea/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
