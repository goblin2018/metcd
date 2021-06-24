package api

import (
	"fmt"
	"strings"

	"github.com/coreos/go-semver/semver"
)

var (
	MinClusterVersion = "3.0.0"
	Version           = "3.6.0-pre"
	APIVersion        = "unknown"

	GitSHA = "Not provided (use ./build instead of go build)"
)

func init() {
	ver, err := semver.NewVersion(Version)
	if err == nil {
		APIVersion = fmt.Sprintf("%d.%d", ver.Major, ver.Minor)
	}
}

type Versions struct {
	Server  string `json:"etcdserver"`
	Cluster string `json:"etcdcluster"`
}

func Cluster(v string) string {
	vs := strings.Split(v, ".")
	if len(vs) <= 2 {
		return v
	}
	return fmt.Sprintf("%s.%s", vs[0], vs[1])
}
