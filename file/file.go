package file

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func GuessExtension(name string) (string, error) {
	list, err := filepath.Glob(name + ".*")
	if err != nil {
		return "", err
	}
	if len(list) != 1 {
		return "", errors.New("none or multiple files with name " + name)
	} else {
		return filepath.Ext(list[0]), nil
	}
}

func NotExist(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return os.IsNotExist(err), err
	}
	defer f.Close()
	return false, errors.New("path already exists")
}

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

func OpenAndTruncate(name string, flags int, perm os.FileMode) (*os.File, error) {
	if (flags & os.O_WRONLY) == 0 {
		return nil, errors.New("write flag is not set")
	}

	f, err := os.OpenFile(name, flags, perm)
	if err != nil {
		return nil, err
	}

	return f, f.Truncate(0)
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