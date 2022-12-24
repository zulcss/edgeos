package image

import (
	"fmt"
	"io/ioutil"
	//"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/zulcss/edgeos/pkg/util"
)

type ImageContext struct {
	image string
}

func Build(arg string) error {
	log.Info("Buildng ", arg)
	// Check docker is installed
	_, err := exec.LookPath("docker")
	if err != nil {
		log.Error("Docker is not installed.")
		return err
	}

	// crete tempdir to build in
	dir, err := ioutil.TempDir("", "stx-build")
	if err != nil {
		return err
	}
	//defer os.RemoveAll(dir)

	_, err = util.SH(fmt.Sprintf("cp -rp images %s", dir))

	/*
		out, err := util.SH("docker images -a")
		if err != nil {
			return err
		}
		fmt.Println(out)
	*/

	return nil
}
