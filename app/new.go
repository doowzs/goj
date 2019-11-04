package app

import (
	"github.com/mholt/archiver"
	"goj/file"
	"io/ioutil"
	"log"
	"os"
)

func runNew(folder string) error {
	res, err := file.IsEmpty(folder)
	if !res {
		return err
	}

	temp, err := ioutil.TempFile(".", "template.*.tar.gz")
	if err != nil {
		return err
	}
	defer os.Remove(temp.Name())

	log.Println("Downloading template archive...")
	url := "https://doowzs.com/goj/template/" + Version + ".tar.gz"
	err = file.DownloadFile(url, temp.Name())
	if err != nil {
		return err
	}

	log.Println("Extracting template archive...")
	return archiver.Unarchive(temp.Name(), folder)
}
