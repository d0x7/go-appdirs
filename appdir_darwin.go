package appdirs

import (
	"os"
	"os/user"
	"path/filepath"
)

type dirs struct {
	name string
}

func (d *dirs) homeDir() string {
	// Check the HOME environment variable first
	if os.Getenv("HOME") != "" {
		return os.Getenv("HOME")
	}

	// otherwise retrieve from the user
	usr, err := user.Current()
	if err != nil {
		return "~"
	}

	return usr.HomeDir
}

func (d *dirs) UserData() string {
	return filepath.Join(d.homeDir(), "Library", "Application Support", d.name)
}

func (d *dirs) UserConfig() string {
	return d.UserData()
}

func (d *dirs) UserCache() string {
	return filepath.Join(d.homeDir(), "Library", "Caches", d.name)
}

func (d *dirs) UserLogs() string {
	return filepath.Join(d.homeDir(), "Library", "Logs", d.name)
}

func (d *dirs) SharedData() string {
	return filepath.Join("/", "Library", "Application Support", d.name)
}

func (d *dirs) SharedConfig() string {
	return d.SharedData()
}

func (d *dirs) SharedCache() string {
	return filepath.Join("/", "Library", "Caches", d.name)
}

func (d *dirs) SharedLogs() string {
	return filepath.Join("/", "Library", "Logs", d.name)
}
