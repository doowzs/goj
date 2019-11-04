package app

import (
	"goj/file"
	"goj/template"
	"os"
)

func runNew(path string) error {
	res, err := file.IsEmpty(path)
	if !res {
		return err
	}

	if os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModeDir)
		if err != nil {
			return err
		}
	}

	return template.Create(path)
}
