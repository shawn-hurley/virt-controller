package settings

import (
	liberr "github.com/konveyor/controller/pkg/error"
	"os"
	"strconv"
)

//
// Global
var Settings = ControllerSettings{}

//
// Settings
type ControllerSettings struct {
	// Roles.
	Role
	// Inventory settings.
	Inventory
	// Migration settings.
	Migration
}

//
// Load settings.
func (r *ControllerSettings) Load() error {
	err := r.Role.Load()
	if err != nil {
		return liberr.Wrap(err)
	}
	err = r.Inventory.Load()
	if err != nil {
		return liberr.Wrap(err)
	}
	err = r.Migration.Load()
	if err != nil {
		return liberr.Wrap(err)
	}

	return nil
}

//
// Get positive integer limit from the environment
// using the specified variable name and default.
func getEnvLimit(name string, def int) (int, error) {
	limit := 0
	if s, found := os.LookupEnv(name); found {
		n, err := strconv.Atoi(s)
		if err != nil {
			return 0, liberr.New(name + " must be an integer")
		}
		if n < 1 {
			return 0, liberr.New(name + " must be >= 1")
		}
		limit = n
	} else {
		limit = def
	}

	return limit, nil
}

//
// Get boolean.
func getEnvBool(name string, def bool) bool {
	boolean := def
	if s, found := os.LookupEnv(name); found {
		parsed, err := strconv.ParseBool(s)
		if err == nil {
			boolean = parsed
		}
	}

	return boolean
}
