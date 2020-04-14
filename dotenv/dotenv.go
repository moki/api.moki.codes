package dotenv

import (
	"bufio"
	"os"
	"strings"
)

// Reader object provides functionality to reads dot env files
type Reader struct {
	filename string
}

// NewReaderT returns new reader
func NewReaderT(filename string) Reader {
	r := Reader{filename}
	return r
}

func errcheck(err error) {
	if err != nil {
		panic(err)
	}
}

// HelloReader displays its settings
func (r Reader) Read() {
	fp, err := os.Open(r.filename)
	errcheck(err)
	defer (func() { errcheck(fp.Close()) })()

	fsc := bufio.NewScanner(fp)
	fsc.Split(bufio.ScanLines)

	for fsc.Scan() {
		pair := strings.Split(fsc.Text(), "=")
		if len(pair) != 2 {
			continue
		}
		errcheck(os.Setenv(pair[0], pair[1]))
	}
}
