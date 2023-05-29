/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/Radical-Egg/notebox/cmd"
	"github.com/Radical-Egg/notebox/internal/database"
)

func main() {
	db, err := database.Connect()

	if err != nil {
		log.Fatalln(err)
	}
	database.CreateNotesTable(db)

	defer db.Close()
	
	cmd.Execute()
}
