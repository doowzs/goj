package app

import (
	"errors"
)

var Version string

func Run(job, src, dst string) error {
	if job[0] == 'n' {
		return runNew(src)
	} else if job[0] == 'g' {
		return runGen(src, dst)
	} else {
		return errors.New("Unsupported job " + job)
	}
}
