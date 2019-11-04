package app

import (
	"errors"
)

var Version string

func Run(job, folder string) error {
	if job[0] == 'n' {
		return runNew(folder)
	} else if job[0] == 'g' {
		return runGen(folder)
	} else {
		return errors.New("Unsupported job " + job)
	}
}
