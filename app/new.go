package app

import (
	"goj/file"
	"goj/template"
)

func runNew(path string) error {
	res, err := file.NotExist(path)
	if !res {
		return err
	}

	return template.Create(path)
}
