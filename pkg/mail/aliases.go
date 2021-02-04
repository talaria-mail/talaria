package mail

import "errors"

// AllowedAlias checks an address again a list of allowed aliases. This can
// be used to check that From address is allowed for a certain user.
func AllowedAlias(from string, allowed []string) error {
	// TODO: Support globs
	for _, alias := range allowed {
		if from == alias {
			return nil
		}
	}
	return errors.New(`No matching alias`)
}
