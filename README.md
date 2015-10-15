# Go bindings for Rancher-metadata

This library is incomplete, but implements a variety of calls against  [rancher-metadata](https://github.com/rancher/rancher-metadata) service

#Example usage

```go
package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/rancher/go-rancher-metadata/metadata"
	"time"
)

const (
	metadataUrl = "http://rancher-metadata"
)

func main() {

  m := metadata.NewHandler(metadataUrl)
  
  version := "init"
  
	for {
		newVersion, err := m.GetVersion()
		if err != nil {
			logrus.Errorf("Error reading metadata version: %v", err)
		} else if version == newVersion {
			logrus.Debug("No changes in metadata version: %s", newVersion)
		} else {
			logrus.Debug("Metadata Version has been changed. Old version: %s. New version: %s.", version, newVersion)
			version = newVersion
		}
		time.Sleep(time.Duration(poll) * time.Millisecond)
	}
}
```
