package app

import (
	"github.com/mholt/archiver"
	"goj/file"
	"io/ioutil"
)

func runNew(src string) error {
	res, err := file.IsEmpty(src)
	if !res {
		return err
	}

	temp, err := ioutil.TempFile(".", "template.*.tar.gz")
	if err != nil {
		return err
	}

	url := "https://doowzs.com/goj/template/" + Version + ".tar.gz"
	err = file.DownloadFile(url, temp.Name())
	if err != nil {
		return err
	}

	return archiver.Unarchive(temp.Name(), src)
}
