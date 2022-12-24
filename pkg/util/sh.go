package util

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func SH(c string) (string, error) {
	log.Info("Running command: ", c)
	cmd := exec.Command("/bin/sh", "-c", c)
	cmd.Env = os.Environ()
	o, err := cmd.CombinedOutput()
	return string(o), err
}
