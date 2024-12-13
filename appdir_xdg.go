//go:build !darwin && !windows
// +build !darwin,!windows

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
	baseDir := os.Getenv("XDG_DATA_HOME")
	if baseDir == "" {
		baseDir = filepath.Join(d.homeDir(), ".local", "share")
	}

	return filepath.Join(baseDir, d.name)
}

func (d *dirs) UserConfig() string {
	baseDir := os.Getenv("XDG_CONFIG_HOME")
	if baseDir == "" {
		baseDir = filepath.Join(d.homeDir(), ".config")
	}

	return filepath.Join(baseDir, d.name)
}

func (d *dirs) UserCache() string {
	baseDir := os.Getenv("XDG_CACHE_HOME")
	if baseDir == "" {
		baseDir = filepath.Join(d.homeDir(), ".cache")
	}

	return filepath.Join(baseDir, d.name)
}

func (d *dirs) UserLogs() string {
	return filepath.Join(d.UserCache(), "logs")
}

func (d *dirs) SharedData() string {
	baseDir := os.Getenv("XDG_DATA_DIRS")
	if baseDir == "" {
		baseDir = filepath.Join("/", "usr", "local", "share")
	}

	return filepath.Join(baseDir, d.name)
}

func (d *dirs) SharedConfig() string {
	baseDir := os.Getenv("XDG_CONFIG_DIRS")
	if baseDir == "" {
		baseDir = filepath.Join("/", "etc", "xdg")
	}

	return filepath.Join(baseDir, d.name)
}

func (d *dirs) SharedCache() string {
	return filepath.Join("/", "tmp", d.name+"-cache")
}

func (d *dirs) SharedLogs() string {
	return filepath.Join(d.SharedData(), "logs")
}
