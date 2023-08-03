package yes

import (
	"os"
)

func Chatter() {

	for {
		os.Stdout.Write([]byte("y\n"))
	}
}
