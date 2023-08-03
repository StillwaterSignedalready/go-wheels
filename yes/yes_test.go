package yes

import (
	"bytes"
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

// TODO: os.Pipe https://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string/10476304#10476304
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
