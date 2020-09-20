package commands

import (
	"os/exec"
	"testing"

	"github.com/profclems/glab/internal/config"

	"github.com/acarl005/stripansi"
	"github.com/stretchr/testify/assert"
)

func Test_AliasSetCmd(t *testing.T) {
	repo := copyTestRepo(t)
	_ = config.DeleteAlias("testmrl")
	var cmd *exec.Cmd

	tests := []struct {
		Name       string
		args       []string
		wantErr    bool
		assertFunc func(t *testing.T, out string)
	}{
		{
			Name:    "Alias name is a command name",
			args:    []string{"mr", "'mr list'"},
			wantErr: true,
			assertFunc: func(t *testing.T, out string) {
				assert.Contains(t, out, "could not create alias: \"mr\" is already a glab command")
			},
		},
		{
			Name: "Is valid",
			args: []string{"testmrl", "mr list"},
			assertFunc: func(t *testing.T, out string) {
				assert.Contains(t, out, "- Adding alias for testmrl: mr list\n✓")
			},
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			cmd = exec.Command(glabBinaryPath, append([]string{"alias", "set"}, test.args...)...)
			cmd.Dir = repo

			b, err := cmd.CombinedOutput()
			if err != nil && !test.wantErr {
				t.Log(string(b))
				t.Fatal(err)
			}
			out := string(b)
			out = stripansi.Strip(out)
			test.assertFunc(t, out)
		})
	}
}