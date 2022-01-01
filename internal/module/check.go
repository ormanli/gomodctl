package module

import (
	"context"
	"errors"
	"sort"

	"github.com/Masterminds/semver"
	"github.com/ormanli/gomodctl/internal"
	"github.com/spf13/viper"
)

// ErrNoVersionAvailable is returned when no version found for a module.
var ErrNoVersionAvailable = errors.New("no version available")

// ErrModuleIgnored is returned when a module is ignored for version check.
var ErrModuleIgnored = errors.New("module ignored")

// Check checks module for updates.
func Check(ctx context.Context, path string) (map[string]internal.CheckResult, error) {
	return getModAndFilter(ctx, path)
}

func getLatestVersion(versions []*semver.Version) (*semver.Version, error) {
	if len(versions) == 0 {
		return nil, ErrNoVersionAvailable
	}

	sort.Sort(semver.Collection(versions))

	lastVersion := versions[len(versions)-1]

	return lastVersion, nil
}

func getModAndFilter(ctx context.Context, path string) (map[string]internal.CheckResult, error) {
	results, err := Parse(ctx, path)
	if err != nil {
		return nil, err
	}

	ignoredModules := getIgnoredModules()

	checkResults := make(map[string]internal.CheckResult)

	for _, result := range results {
		checkResult := internal.CheckResult{
			LocalVersion: result.LocalVersion,
		}

		_, isIgnored := ignoredModules[result.Path]
		if isIgnored {
			checkResult.Error = ErrModuleIgnored
		} else {
			latestVersion, err := getLatestVersion(result.AvailableVersions)

			if err != nil {
				checkResult.Error = err
			}

			if latestVersion != nil {
				checkResult.LatestVersion = latestVersion
			}
		}

		checkResults[result.Path] = checkResult
	}

	return checkResults, nil
}

type void struct{}

var member void

func getIgnoredModules() map[string]void {
	s := make(map[string]void)

	im := viper.GetStringSlice("ignored_modules")
	for _, m := range im {
		s[m] = member
	}

	return s
}
