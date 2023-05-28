/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/Radical-Egg/notebox/cmd"
	"github.com/Radical-Egg/notebox/internal/database"
)

func main() {
	database.Init()
	cmd.Execute()
}
