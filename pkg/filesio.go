package pkg

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadAndDisplayTextFile(filepath string, w io.Writer, showLineCounts bool) error {
	file, err := os.Open(filepath)

	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCounter := 1

	for scanner.Scan() {
		lineCounterDisplay := ""

		if showLineCounts {
			lineCounterDisplay = fmt.Sprintf("%d:", lineCounter)
			lineCounter++
		}

		fmt.Fprintf(w, "%s %s\n", lineCounterDisplay, scanner.Text())
	}

	return nil
}
func ReadAndDisplayStandardInput() error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return scanner.Err()
}
