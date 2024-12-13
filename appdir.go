// Package appdirs gets application directories such as config and cache.
package appdirs

// Dirs requests application directories paths.
type Dirs interface {
	homeDir() string

	// UserData returns user-specific data directory.
	UserData() string
	// UserConfig returns the user-specific config directory.
	UserConfig() string
	// UserCache returns the user-specific cache directory.
	UserCache() string
	// UserLogs returns user-specific logs directory.
	UserLogs() string

	// SharedData returns the site-specific data directory.
	SharedData() string
	// SharedConfig returns the site-specific config directory.
	SharedConfig() string
	// SharedCache returns the site-specific cache directory.
	SharedCache() string
	// SharedLogs returns the site-specific logs directory.
	SharedLogs() string
}

// New creates a new App with the provided name.
func New(name string) Dirs {
	return &dirs{name: name}
}
