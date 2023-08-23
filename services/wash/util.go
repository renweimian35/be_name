package wash

import (
	"encoding/json"
	"os"
)

func resolve(path string, info any) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	return json.Unmarshal(content, info)
}
