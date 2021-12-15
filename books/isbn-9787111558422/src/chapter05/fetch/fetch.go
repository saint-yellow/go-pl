package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (string, int64, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}

	defer response.Body.Close()

	local := path.Base(response.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err := io.Copy(f, response.Body)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}

func main() {
	fmt.Println(fetch(os.Args[1]))
}