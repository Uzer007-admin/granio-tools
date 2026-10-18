package commands

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/mod/modfile"
)

func TestEnsureGranioDependency(t *testing.T) {
	for _, selected := range []string{"", "v2.8.0-custom.7"} {
		t.Run("selected="+selected, func(t *testing.T) {
			dir := t.TempDir()
			path := filepath.Join(dir, "go.mod")
			content := "module example.test/app\n\ngo 1.27.1\n"
			if selected != "" {
				content += "\nrequire " + fyneModule + " " + selected + "\n"
			}
			require.NoError(t, os.WriteFile(path, []byte(content), 0600))
			require.NoError(t, ensureGranioDependency(dir, false))
			data, err := os.ReadFile(path)
			require.NoError(t, err)
			mod, err := modfile.Parse(path, data, nil)
			require.NoError(t, err)
			require.Len(t, mod.Require, 1)
			require.Equal(t, fyneModule, mod.Require[0].Mod.Path)
			if selected != "" {
				require.Equal(t, content, string(data), "выбранная приложением версия сохраняется")
			} else {
				require.Equal(t, granioVersionForInit(), mod.Require[0].Mod.Version)
			}
		})
	}
}
