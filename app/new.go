package app

import (
	"github.com/mholt/archiver"
	"goj/file"
	"io/ioutil"
	"os"
)

func runNew(src string) error {
	res, err := file.IsEmpty(src)
	if !res {
		return err
	}

	temp, err := ioutil.TempFile(".", "template.*.tar.xz")
	if err != nil {
		return err
	}
	defer os.Remove(temp.Name())

	url := "https://doowzs.com/goj/template/" + Version + ".tar.xz"
	err = file.DownloadFile(url, temp.Name())
	if err != nil {
		return err
	}

	return archiver.Unarchive(temp.Name(), src)
}
