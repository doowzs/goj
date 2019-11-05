package app

import (
	"goj/file"
	"goj/template"
	"os"
)

func runGen(path string) error {
	f, err := file.OpenAndTruncate(path + "/dist.xml", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return template.Generate(f, path)
}