// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package lockfiles

import (
	"strings"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestImageLockNonDigestUnmarshalError(t *testing.T) {
	imageLockYaml := []byte(`apiVersion: imgpkg.carvel.dev/v1alpha1
kind: ImagesLock
spec:
  images:
  - image: nginx:v1`)

	var imageLock ImageLock
	err := yaml.Unmarshal(imageLockYaml, &imageLock)

	if err == nil {
		t.Fatalf("Expected unmarshal to error")
	}

	if msg := err.Error(); !(strings.Contains(msg, "to be in digest form") && strings.Contains(msg, "nginx:v1")) {
		t.Fatalf("Expected unmarshal to fail due to tag ref in lock file")
	}
}
