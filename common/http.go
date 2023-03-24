package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (app *App) ParseJSON(r *http.Request, dst interface{}) error {
	b, err := io.ReadAll(r.Body)
	if r.Body != nil {
		defer r.Body.Close()
	}
	if err != nil {
		return err
	}
	if dst != nil {
		if err := json.Unmarshal(b, dst); err != nil {
			return err
		}
		return nil
	}
	return errors.New("no destination for ParseJSON")
}

func (app *App) GetJSON(url string, dst interface{}) error {
	resp, err := app.httpClient.Get(url)
	if err != nil {
		return err
	}
	b, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	if dst != nil {
		if err := json.Unmarshal(b, dst); err != nil {
			return err
		}
	}
	return nil
}

func (app *App) PostJSON(url string, src, dst interface{}, expectingStatus int, headers ...map[string]string) error {

	method := "POST"

	var buf *bytes.Buffer
	if src != nil {
		b, err := app.MarshalJSON(src)
		if err != nil {
			return err
		}
		buf = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			req.Header.Set(k, v)
		}
		if app.IsDebug() {
			for k, v := range req.Header {
				println(method, url, k, v)
			}
		}
	}

	resp, err := app.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != expectingStatus {
		s := fmt.Sprintf("invalid status code for http post, expecting %d, got %d", expectingStatus, resp.StatusCode)
		return errors.New(s)
	}

	b, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	if dst != nil {
		if err := json.Unmarshal(b, dst); err != nil {
			return err
		}
	}
	return nil
}
