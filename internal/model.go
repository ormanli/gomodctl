package internal

import (
	"github.com/Masterminds/semver"
)

// CheckResult is exported.
type CheckResult struct {
	LocalVersion  *semver.Version
	LatestVersion *semver.Version
	Error         error
}
