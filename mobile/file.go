package mobile

import (
	"os"

	"github.com/gopub/log"
)

func MustMkdir(dir string) {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatalf("Make dir: %s, %v", dir, err)
	}
}
