package common

import (
	"encoding/json"
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

	// make sure we ignore subordinate drivel
	j := content[in : out+1]

	return json.Unmarshal([]byte(j), dst)
}
