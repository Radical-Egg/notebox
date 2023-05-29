/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/Radical-Egg/notebox/internal/database"
	"github.com/Radical-Egg/notebox/internal/editor"
	"github.com/spf13/cobra"
	"github.com/tjarratt/babble"
)

// writeCmd represents the write command
var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Use this function to write a new note",
	Long: `This function can be used to write a new note. It leverages the
	$EDITOR variable.`,
	Run: func(cmd *cobra.Command, args []string) {
		var title string

		if len(args) == 0 {
			babbler := babble.NewBabbler()
			babbler.Separator = " "
			title = babbler.Babble()
		} else {
			title = args[0]
		}

		ed := editor.Editor{}

		ed.Open()

		note := database.Note{Title: title, Content: ed.Content}
		
		db, err := database.Connect()

		if err != nil {
			log.Fatalln(err)
		}

		defer db.Close()
		
		t := database.CreateNote(db, note)

		if t != nil {
			log.Fatalln(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(writeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// writeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// writeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
