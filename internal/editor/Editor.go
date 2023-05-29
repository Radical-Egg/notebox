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

func(editor *Editor) Open() error {
	var wg sync.WaitGroup

	done := make(chan struct{})

	wg.Add(1)
	
	go func(ed *Editor) error {
		defer wg.Done()
		tempFile, err := os.CreateTemp("", "prefix")

		if err != nil {
			return err
		}

		defer tempFile.Close()
		defer os.Remove(tempFile.Name())

		if ed.Content != "" {
			_, err = tempFile.Write([]byte(ed.Content))
			if err != nil {
				return err
			}
		}

		editor := os.Getenv("EDITOR")

		if editor == "" {
			return fmt.Errorf("cannot open $EDITOR, value %s", editor)
		}

	
		write := exec.Command(editor, tempFile.Name())
		write.Stdout = os.Stdout
		write.Stdin = os.Stdin
		write.Stderr = os.Stderr
		
		if err := write.Run(); err != nil {
			return err
		}

		content, err := os.ReadFile(tempFile.Name())

		if err != nil {
			return err
		}

		ed.Content = string(content)
		done <- struct{}{}
		return nil
	} (editor)

	<-done

	wg.Wait()

	return nil
}