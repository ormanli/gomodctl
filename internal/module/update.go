package module

import (
	"context"
	"io/ioutil"
	"path/filepath"

	"github.com/ormanli/gomodctl/internal"
	"golang.org/x/mod/modfile"
)

const (
	goMod       = "go.mod"
	goModBackup = "go.mod.backup"
)

// Update is exported
func Update(ctx context.Context, path string) (map[string]internal.CheckResult, error) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	absoluteFile := filepath.Join(absolutePath, goMod)
	backupFile := filepath.Join(absolutePath, goModBackup)

	content, err := ioutil.ReadFile(absoluteFile)
	if err != nil {
		return nil, err
	}

	parse, err := modfile.Parse("go.mod", content, nil)
	if err != nil {
		return nil, err
	}

	latestMinors, err := getModAndFilter(ctx, absolutePath)
	if err != nil {
		return nil, err
	}

	updates := 0
	requires := parse.Require

	for i := range requires {
		require := requires[i]
		if !require.Indirect {
			if result, ok := latestMinors[require.Mod.Path]; ok && result.Error == nil && result.LatestVersion.GreaterThan(result.LocalVersion) {
				requires[i].Mod.Version = result.LatestVersion.Original()

				updates++
			}
		}
	}

	if updates > 0 {
		parse.SetRequire(requires)
		parse.Cleanup()
		//parse.SortBlocks()

		format, err := parse.Format()
		if err != nil {
			return nil, err
		}

		err = ioutil.WriteFile(backupFile, content, 0666)
		if err != nil {
			return nil, err
		}

		err = ioutil.WriteFile(absoluteFile, format, 0666)
		if err != nil {
			return nil, err
		}
	}

	return latestMinors, nil
}
