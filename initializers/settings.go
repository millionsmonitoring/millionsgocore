package initializers

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/millionsmonitoring/millionsgocore/env"
	"github.com/millionsmonitoring/millionsgocore/helpers"
)

// initializer for loading configs/settings.yml file
// this provides a way to load configuration based on the env
// this will configurations that are needed for the application to run
// but are not something if leaked would cause an issue
func LoadSettings[T helpers.Config](ctx context.Context) (T, error) {
	loadedSettings, err := helpers.CheckOrParseConfig[T](fmt.Sprintf("%s.yml", env.Env()))
	if err != nil {
		slog.ErrorContext(ctx, "unable to load settings", "err", err)
		return loadedSettings, err
	}

	return loadedSettings, nil
}
