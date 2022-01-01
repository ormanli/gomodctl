package module

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/spf13/viper"
)

var regex = regexp.MustCompile(`({([^}]*)})`)

var go115 = semver.MustParse("1.15.0")

type item struct {
	Path     string   `json:"Path"`
	Version  string   `json:"Version"`
	Versions []string `json:"Versions"`
	Indirect bool     `json:"Indirect"`
	Main     bool     `json:"Main"`
	Dir      string   `json:"Dir"`
	GoMod    string   `json:"GoMod"`
}

// PackageResult contains module specific information.
type PackageResult struct {
	Path              string
	LocalVersion      *semver.Version
	AvailableVersions []*semver.Version
	Dir               string
}

// Parse parses go.mod.
func Parse(ctx context.Context, path string) ([]PackageResult, error) {
	goVersion, err := goRuntimeVersion(ctx)
	if err != nil {
		return nil, err
	}

	args := []string{"list", "-m", "-versions", "-json", "-mod=mod", "all"}
	if goVersion.LessThan(go115) {
		args = []string{"list", "-m", "-versions", "-json", "all"}
	}

	cmd := exec.CommandContext(ctx, "go", args...)

	if path != "" {
		home := viper.GetString("home")

		if strings.HasPrefix(path, home) {
			l := path[len(home):]
			cmd.Dir = filepath.Join(home, l)
		} else {
			cmd.Dir = filepath.Join(home, path)
		}
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		if len(out) > 0 {
			return nil, fmt.Errorf("with output [%s] %w", out, err)
		}

		return nil, err
	}

	versionOutputs := regex.FindAll(out, -1)

	var result []PackageResult

	for _, versionOutput := range versionOutputs {
		it := item{}

		err := json.Unmarshal(versionOutput, &it)
		if err != nil {
			return nil, err
		}

		if !it.Indirect && !it.Main {
			availableVersions := make([]*semver.Version, len(it.Versions))

			for i, version := range it.Versions {
				availableVersions[i] = semver.MustParse(version)
			}

			result = append(result, PackageResult{
				Path:              it.Path,
				LocalVersion:      semver.MustParse(it.Version),
				Dir:               it.Dir,
				AvailableVersions: availableVersions,
			})
		}
	}

	return result, nil
}

func goRuntimeVersion(ctx context.Context) (*semver.Version, error) {
	cmd := exec.CommandContext(ctx, "go", "version")

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	r := regexp.MustCompile(`(go version go)(.*)( .+)`)
	find := r.FindSubmatch(out)

	version, err := semver.NewVersion(string(find[2]))
	if err != nil {
		return nil, err
	}

	return version, nil
}
