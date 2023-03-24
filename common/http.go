package common

import (
	"bytes"
	"encoding/json"
	"errors"
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

func (app *App) PostJSON(url string, src, dst interface{}, headers ...map[string]string) error {

	var buf *bytes.Buffer
	if src != nil {
		b, err := json.Marshal(src)
		if err != nil {
			return err
		}
		buf = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(url, "application/json", buf)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			req.Header.Set(k, v)
		}
	}

	resp, err := app.httpClient.Do(req)
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
