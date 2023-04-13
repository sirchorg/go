package cloudfunc

import (
	"encoding/json"
	"os"

	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetSecretFromVolume(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func HandleCORS(w http.ResponseWriter, r *http.Request, origin string) bool {
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Cache-Control", "no-store")
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.WriteHeader(http.StatusNoContent)
		return true
	}
	return false
}

func HttpError(w http.ResponseWriter, err error, status int) {
	if err != nil {
		log.Println(err)
	}
	http.Error(w, err.Error(), status)
	w.Write([]byte(fmt.Sprintf("REQUEST FAILED: %d %s", status, err)))
}

func ParseJSON(r *http.Request, dst interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	if r.Body != nil {
		r.Body.Close()
	}
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, dst); err != nil {
		return err
	}
	return nil
}

func ServeJSON(w http.ResponseWriter, src interface{}) error {
	b, err := json.Marshal(src)
	if err != nil {
		return err
	}
	w.Write(b)
	return nil
}

func ParamValue(r *http.Request, pos int) string {
	return strings.Split(r.URL.Path, "/")[pos]
}
