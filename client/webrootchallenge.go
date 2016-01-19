package client

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/juju/errors"
	log "github.com/sirupsen/logrus"
)

type WebRootChallengeResponder struct {
	rootPath string
}

func NewWebRootChallengeResponder(rootPath string) (*WebRootChallengeResponder, error) {
	cr := WebRootChallengeResponder{rootPath: rootPath}
	return &cr, nil
}

func (h *WebRootChallengeResponder) SetResource(p, resource string) error {
	log.WithFields(log.Fields{
		"mode":     "webroot",
		"path":     p,
		"resource": resource,
	}).Debugln("SetResource")

	if err := os.MkdirAll(h.rootPath, 0755); err != nil {
		return errors.Annotate(err, "ensure root path exists")
	}

	resPath := path.Join(h.rootPath, p)
	if err := ioutil.WriteFile(resPath, []byte(resource), 0755); err != nil {
		return errors.Annotate(err, "write ressource")
	}
	return nil
}
