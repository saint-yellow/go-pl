package log

import (
	"bytes"
	"fmt"
	stdlog "log"
	"net/http"

	"github.com/saint-yellow/go-pl/toys/distributed-system/registry"
)

func SetClientLogger(serviceURL string, clientService registry.ServiceName) {
	stdlog.SetPrefix(fmt.Sprintf("[%v] - ", clientService))
	stdlog.SetFlags(0)
	stdlog.SetOutput(&clientLogger{url: serviceURL})
}

type clientLogger struct {
	url string
}

func (cl clientLogger) Write(data []byte) (int, error) {
	var err error = nil

	buf := bytes.NewBuffer(data)
	res, err := http.Post(cl.url + "/log", "text/plain", buf)
	if err != nil {
		return 0, err
	}

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("failed to send log message, service responded with status code %v", res.StatusCode)
		return 0, err
	}

	return len(data), nil
}