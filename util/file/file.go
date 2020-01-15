package file

import (
	"io/ioutil"
	"os"
	"time"
)

// WriteFile
func WriteFile(fileName string, content string) {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err = f.Write([]byte(content))
	if err != nil {
		panic(err)
	}
}

// ReadFile
func ReadFile(fileName string) (string, bool) {
	f, err := os.OpenFile(fileName, os.O_RDONLY, 0600)
	if err != nil {
		return "", false
	}

	defer f.Close()

	s, err := f.Stat()
	if err != nil {
		panic(err)
	}

	modTime := s.ModTime()
	now := time.Now()
	// mod time > 1d, file expire
	if now.Unix()-modTime.Unix() > 24*60*60*1000 {
		return "", false
	}

	contentByte, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return string(contentByte), true
}
