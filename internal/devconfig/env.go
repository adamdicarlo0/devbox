package devconfig

import (
	"context"
	"fmt"
	"os"

	"go.jetpack.io/devbox/internal/boxcli/usererr"
	"go.jetpack.io/devbox/internal/integrations/envsec"
)

func (c *Config) ComputedEnv(
	ctx context.Context,
	projectDir string,
) (map[string]string, error) {
	env := map[string]string{}
	var err error
	if c.IsEnvsecEnabled() {
		env, err = envsec.Env(ctx, projectDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading secrets from envsec: %s\n\n", err)
		}
	} else if c.EnvFrom != "" {
		return nil, usererr.New("unknown from_env value: %s", c.EnvFrom)
	}
	for k, v := range c.Env {
		env[k] = v
	}
	return env, nil
}

func (c *Config) IsEnvsecEnabled() bool {
	return c.EnvFrom == "envsec"
}
