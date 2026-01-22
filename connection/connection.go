package connection

import (
	"GoTTP/http"
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

func ReadAndParseRequest(r *bufio.Reader) (*http.Request, error) {

	line, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}

	line = strings.TrimRight(line, "\r\n")
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return nil, errors.New("invalid request line")
	}

	req := &http.Request{
		Method:  parts[0],
		Path:    parts[1],
		Version: parts[2],
		Headers: make(map[string]string),
	}
	contentLength := 0
	for {
		h, err := r.ReadString('\n')
		if err != nil {
			return nil, err
		}

		h = strings.TrimRight(h, "\r\n")
		if h == "" {
			break // end of headers
		}

		kv := strings.SplitN(h, ":", 2)
		if len(kv) != 2 {
			continue
		}

		key := strings.TrimSpace(kv[0])
		val := strings.TrimSpace(kv[1])

		req.Headers[key] = val

		if strings.EqualFold(key, "Content-Length") {
			contentLength, _ = strconv.Atoi(val)
		}
		if strings.EqualFold(key, "Connection") &&
			strings.EqualFold(val, "close") {
			req.Close = true
		}
	}

	// ---- Body ----
	if contentLength > 0 {
		req.Body = make([]byte, contentLength)
		_, err := io.ReadFull(r, req.Body)
		if err != nil {
			return nil, err
		}
	}

	return req, nil
}
