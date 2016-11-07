package metadata

import (
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
)

func (m *client) OnChange(intervalSeconds int, do func(string)) {
	interval := time.Duration(intervalSeconds)
	version := "init"

	for {
		newVersion, err := m.waitVersion(intervalSeconds, version)
		if err != nil {
			logrus.Errorf("Error reading metadata version: %v", err)
			time.Sleep(interval * time.Second)
		} else if version == newVersion {
			logrus.Debug("No changes in metadata version")
			time.Sleep(interval * time.Second)
		} else {
			logrus.Debugf("Metadata Version has been changed. Old version: %s. New version: %s.", version, newVersion)
			version = newVersion
			do(newVersion)
		}
	}
}

func (m *client) waitVersion(maxWait int, version string) (string, error) {
	resp, err := m.SendRequest(fmt.Sprintf("/version?wait=true&value=%s&maxWait=%d"))
	if err != nil {
		return "", err
	}
	return string(resp[:]), nil
}
