package file

import (
	"errors"
	"io"
	"net/http"
	"os"
)

func IsEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return os.IsNotExist(err), err
	}
	defer f.Close()

	_, err = f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, errors.New("folder not empty")
}

func CopyFolder(src, dst string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	files, err := f.Readdir(0)
	for idx := range(files) {
		err = CopyFile(src + "/" + files[idx].Name(), dst + "/" + files[idx].Name())
		if err != nil {
			return err
		}
	}
	return f.Close()
}

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	// use CopyN to avoid memory explosion
	_, err = io.CopyN(out, in, 1024)
	if err != io.EOF {
		return err
	}
	return out.Close()
}

func DownloadFile(url, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.CopyN(out, resp.Body, 1024)
	if err != io.EOF {
		return err
	}
	return nil
}