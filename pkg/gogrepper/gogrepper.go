package gogrepper

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sync"

	"github.com/fatih/color"
)

var red = color.New(color.FgRed).Add(color.Bold).SprintFunc()

func IsInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func PipeReader(r io.Reader, w io.Writer, pattern string) error {
	scanner := bufio.NewScanner(bufio.NewReader(r))
	processLineConc(scanner, pattern)
	return nil
}

func ReadFile(path string, pattern string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	processLineConc(scanner, pattern)

	return nil
}

func matchText(jobs <-chan string, wg *sync.WaitGroup, pattern string) {
	defer wg.Done()

	var regexer = regexp.MustCompile(pattern)

	for line := range jobs {
		// Replace all matches with a red coloring
		match := regexer.MatchString(line)
		if match {
			formattedLine := regexer.ReplaceAllString(line, red("$0"))
			fmt.Println(formattedLine)
		}

	}
}

func processLineConc(scanner *bufio.Scanner, pattern string) error {

	// Channel to hold lines
	job := make(chan string)

	go func() {
		for scanner.Scan() {
			job <- scanner.Text()
		}
		close(job)
	}()

	wg := new(sync.WaitGroup)
	for w := 1; w <= 100; w++ {
		wg.Add(1)
		go matchText(job, wg, pattern)
	}
	wg.Wait()

	return nil
}
