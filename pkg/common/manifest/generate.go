package manifest

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/alecthomas/chroma/quick"
)

func Generate(v any, raw bool) error {
	out, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	if raw {
		_, err = fmt.Println(string(out))
		return err
	}

	return quick.Highlight(os.Stdout, string(out), "json", "terminal", "monokai")
}
