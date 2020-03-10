package compile

import (
	"errors"
  "fmt"
	"goj/file"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

func Compile(path, name, ext string) (string, error) {
	var (
		f   *os.File
		cmd *exec.Cmd
		exe string
		err error
	)

	tmp := path + "tmp/"
	notExist, _ := file.NotExist(tmp)
	if notExist {
		err = os.Mkdir(tmp, os.ModeDir|0755)
		if err != nil {
			return "", err
		}
	}

	if runtime.GOOS == "windows" {
		f, err = ioutil.TempFile(tmp, name + ".*.exe")
	} else {
		f, err = ioutil.TempFile(tmp, name + ".*.out")
	}
	if err != nil {
		return "", err
	}

	exe = f.Name()
	err = f.Close()
	if err != nil {
		return "", err
	}

	switch ext {
	case ".c":
		cmd = exec.Command("gcc", "-fno-asm", "-Wall", "-lm", "-O2",
      "-std=c99", "-DONLINE_JUDGE", "-o", exe, path + name + ext)
		break
	case ".cc":
	case ".cpp":
		cmd = exec.Command("g++", "-fno-asm", "-Wall", "-lm", "-O2",
			"-std=c++14", "-DONLINE_JUDGE", "-o", exe, path + name + ext)
		break
	default:
		return "", errors.New("unspported source language")
	}

  output, err := cmd.CombinedOutput()
  if err != nil {
    fmt.Println(fmt.Sprint(err) + ": " + string(output))
    return "", err
  }
  return exe, nil
}
