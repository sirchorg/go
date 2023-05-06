package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func (app *App) ParseContentForObjectOrArrayJSON(content string, dst interface{}) error {

	// trim the content to only raw json
	inO := strings.Index(content, "{")
	inA := strings.Index(content, "[")
	outO := strings.LastIndex(content, "}")
	outA := strings.LastIndex(content, "]")

	if app.debugMode {
		println(">>", inO, string(content[inO]), inA, string(content[inA]))
		println(">>", outO, string(content[outO]), outA, string(content[outA]))
	}

	if (inO < 0 || outO < 0) && (inA < 0 || outA < 0) {
		return fmt.Errorf("out of bounds error, no array or object? %d, %d, %d, %d", inO, inA, outO, outA)
	}

	var in, out int
	if inO > 0 && inO > inA {
		in = inO
		out = outO
		if app.debugMode {
			println("receiving object")
		}
	} else {
		in = inA
		out = outA
		if app.debugMode {
			println("receiving array")
		}
	}

	// make sure we ignore subordinate drivel (safely)
	if out >= len(content) {
		return errors.New("preventing panid when trimming content")
	}
	j := content[in : out+1]

	if app.debugMode {
		println("DEBUG", j)
	}

	return json.Unmarshal([]byte(j), dst)
}
