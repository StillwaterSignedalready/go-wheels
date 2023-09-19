package yes

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func every[T comparable](s []T, f func(e T) bool) bool {
	result := true
	for _, e := range s {

		result = result && f(e)
		if !result {
			break
		}
	}
	return result
}

func TestChatter(t *testing.T) {
	// TODO: race condition?
	var buf bytes.Buffer
	go func() {
		Chatter(&buf)
	}()
	time.Sleep(time.Second)

	lines := strings.Split(buf.String(), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	assert.True(t, every(lines, func(e string) bool { return e == "y" }))
}

func printFlood() {
	for i := 0; i < 64<<10; i++ {
		fmt.Println("y")
	}
}

// https://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string/10476304#10476304
func TestChatterByPipe(t *testing.T) {
	/*
	* Chatter write string in pipe
	* get string by pipe
	* check string
	 */
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
		return
	}
	originalStdout := os.Stdout
	os.Stdout = w
	defer func() {
		os.Stdout = originalStdout
	}()

	stdStrCh := make(chan string)
	go func() { // get string from pipe after input over or pipe's limit reached
		bytes, _ := io.ReadAll(r)
		stdStrCh <- string(bytes)
	}()

	// code to test
	// pipe's buffer could be filled and goroutine blocked
	go func() {
		Chatter(os.Stdout)
		// printFlood()
	}()
	time.Sleep(time.Second)

	w.Close() // must before <-stdStrCh, prevent dead lock
	stdStr := <-stdStrCh

	lines := strings.Split(stdStr, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	assert.True(t, every(lines, func(e string) bool { return e == "y" }))
}
