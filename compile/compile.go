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

func Compile(path, name, ext string) (string, string, []string, error) {
	var (
		f         *os.File
		cmd       *exec.Cmd
		tempfile  string
		runnable  string
		arguments []string
		err       error
	)

	tempFolder := path + "tmp/"
	notExist, _ := file.NotExist(tempFolder)
	if notExist {
		err = os.Mkdir(tempFolder, os.ModeDir|0755)
		if err != nil {
			return "", "", nil, err
		}
	}

	if ext == ".java" {
		tempfile = tempFolder + name + ".class"
		runnable = "java"
		arguments = []string{"-classpath", tempFolder, name}
	} else {
		if runtime.GOOS == "windows" {
			f, err = ioutil.TempFile(tempFolder, name+".*.exe")
		} else {
			f, err = ioutil.TempFile(tempFolder, name+".*.out")
		}
		if err != nil {
			return "", "", nil, err
		}

		tempfile = f.Name()
		runnable = f.Name()
		arguments = []string{}
		err = f.Close()
		if err != nil {
			return "", "", nil, err
		}
	}

	switch ext {
	case ".c":
		cmd = exec.Command("gcc", "-fno-asm", "-Wall", "-lm", "-O2",
			"-std=c99", "-DONLINE_JUDGE", "-o", tempfile, path+name+ext)
		break
	case ".cc":
	case ".cpp":
		cmd = exec.Command("g++", "-fno-asm", "-Wall", "-lm", "-O2",
			"-std=c++14", "-DONLINE_JUDGE", "-o", tempfile, path+name+ext)
		break
	case ".java":
		cmd = exec.Command("javac", "-J-Xms32m", "-J-Xmx256m", "-d", tempFolder, path+name+ext)
	default:
		return "", "", nil, errors.New("unspported source language")
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		return "", "", nil, err
	}
	return tempfile, runnable, arguments, nil
}
