package app

import (
	"errors"
)

var Version string

func Run(job, path string) error {
	if job[0] == 'n' {
		return runNew(path)
	} else if job[0] == 'g' {
		return runGen(path)
	} else {
		return errors.New("Unsupported job " + job)
	}
}
