package xkcd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

const (
	BaseURL = "https://xkcd.com"
	DefaultClientTimeout = 30 * time.Second
	LatestComic ComicNumber = 0
)

type XKCDClient struct {
	client *http.Client
	baseURL string
}

func NewXKCDClient() *XKCDClient {
	return &XKCDClient{
		client: &http.Client{
			Timeout: DefaultClientTimeout,
		},
		baseURL: BaseURL,
	}
}

func (xc *XKCDClient) SetTimeout(d time.Duration) {
	xc.client.Timeout = d
}

func (xc *XKCDClient) buildURL(n ComicNumber) string {
	if n == LatestComic {
		return fmt.Sprintf("%s/info.0.json", xc.baseURL)
	} else {
		return fmt.Sprintf("%s/%d/info.0.json",xc.baseURL, n)
	}
}

func (xc *XKCDClient) Download(url, savePath string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	absSavePath, _ := filepath.Abs(savePath)
	filePath := fmt.Sprintf("%s/%s", absSavePath, path.Base(url))

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func (xc *XKCDClient) Fetch(n ComicNumber, downloadPath string) (Comic, error) {
	response, err := xc.client.Get(xc.buildURL(n))
	if err != nil {
		return Comic{}, nil
	}

	defer response.Body.Close()

	var cr ComicResponse
	if err := json.NewDecoder(response.Body).Decode(&cr); err != nil {
		return Comic{}, err
	}

	if downloadPath != "" {
		if err := xc.Download(cr.Img, downloadPath); err != nil {
			fmt.Printf("Failed to download image: %s\n", err)
			return Comic{}, err
		}
	}

	return cr.Comic(), nil
}