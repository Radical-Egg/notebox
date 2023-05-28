package editor

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

type Editor struct {
	Content	string
}

func(editor *Editor) Open() {
	var wg sync.WaitGroup

	done := make(chan struct{})


	wg.Add(1)
	
	go func(ed *Editor) {
		defer wg.Done()
		tempFile, err := os.CreateTemp("", "prefix")
		if err != nil {
			fmt.Println("Error creating temporary file:", err)
			return
		}

		defer tempFile.Close()
		defer os.Remove(tempFile.Name())
	
		editor := os.Getenv("EDITOR")
	
		// return error here
		if editor == "" {
			return
		}
	
		write := exec.Command(editor, tempFile.Name())
		write.Stdout = os.Stdout
		write.Stdin = os.Stdin
		write.Stderr = os.Stderr
		
		if err := write.Run(); err != nil {
			// return err
			return
		}

		content, err := os.ReadFile(tempFile.Name())

		if err != nil {
			return
		}

		ed.Content = string(content)
		done <- struct{}{}
	} (editor)

	<-done

	wg.Wait()
}