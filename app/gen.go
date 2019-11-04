package app

import (
	"goj/template"
	"os"
)

func runGen(path string) error {
	f, err := os.OpenFile(path + "/dist.xml", os.O_WRONLY, 0644)
	if os.IsNotExist(err) {
		f, err = os.Create(path + "/dist.xml")
	}
	if err != nil {
		return err
	}
	defer f.Close()

	err = f.Truncate(0)
	if err != nil {
		return err
	}

	return template.Generate(f, path)
}