/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Radical-Egg/notebox/internal/database"
	"github.com/Radical-Egg/notebox/internal/editor"
	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open an existing note",
	Long: `This command will
	open an existing note. If not argument is specified it will
	open the last note that was editted`,
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("Invalid argument: %s\n", arg)
			return
		}
	
		db, err := database.Connect()

		if err != nil {
			log.Fatalln(err)
		}

		defer db.Close()

		query, query_err := database.GetNoteById(db, num)

		if query_err != nil {
			log.Fatalln(query_err)
		}

		ed := editor.Editor{Content: query.Content}

		ed.Open()

		note := database.Note{Title: query.Title, Content: ed.Content, ID: query.ID}

		update := database.UpdateNote(db, note)

		if update != nil {
			log.Fatalln(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
