package main

import (
	"encoding/json"
	"os"
	"slices"
)

func writeJSONReport(pages map[string]PageData, filename string) error {
	output := []PageData{}
	keys := make([]string, 0, len(pages))
	for k := range pages {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, k := range keys {
		output = append(output, pages[k])
	}
	data, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return err
	}
	os.WriteFile(filename, data, 0644)

	return nil
}
