package image

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/zulcss/edgeos/pkg/util"
)

func Build(BaseDir string) error {
	log.Info("Image pre-flight checks")
	// Check docker is installed
	log.Info("Checking for docker executable...")
	_, err := exec.LookPath("docker")
	if err != nil {
		log.Error("Docker is not installed.")
		return err
	}

	// crete tempdir to build in
	log.Info("Creating directory to build docker images")
	dir, err := ioutil.TempDir("", "stx-build")
	if err != nil {
		return err
	}
	log.Info("Created ", dir)
	defer os.RemoveAll(dir)

	_, err = util.SH(fmt.Sprintf("cp -rp %s/base/* %s", BaseDir, dir))
	if err != nil {
		log.Error("Failed to copy: ", err)
		return err
	}

	_, err = util.SH(fmt.Sprintf("docker build -t stx-image %s", dir))
	if err != nil {
		log.Error("Failed to ", err)
		return err
	}
	return nil
}
