//go:build darwin
// +build darwin

package appdirs

import (
	"os"
	"testing"
)

var app Dirs

func init() {
	if err := os.Setenv("HOME", "/home/user"); err != nil {
		panic(err)
	}
	app = New("TestApplication")
}

func TestUserData(t *testing.T) {
	expected := "/home/user/Library/Application Support/TestApplication"
	result := app.UserData()
	if result != expected {
		t.Errorf("expected %q, found %q", expected, result)
	}
}

func TestUserConfig(t *testing.T) {
	expected := "/home/user/Library/Application Support/TestApplication"
	result := app.UserConfig()
	if result != expected {
		t.Errorf("expected %q, found %q", expected, result)
	}
}

func TestUserCache(t *testing.T) {
	expected := "/home/user/Library/Caches/TestApplication"
	result := app.UserCache()
	if result != expected {
		t.Errorf("expected %q, found %q", expected, result)
	}
}

func TestUserLogs(t *testing.T) {
	expected := "/home/user/Library/Logs/TestApplication"
	result := app.UserLogs()
	if result != expected {
		t.Errorf("expected %q, found %q", expected, result)
	}
}

func TestSharedData(t *testing.T) {
	expected := "/Library/Application Support/TestApplication"
	result := app.SharedData()
	if result != expected {
		t.Errorf("expected %q, found %q", expected, result)
	}
}

func TestSharedConfig(t *testing.T) {
	expected := "/Library/Application Support/TestApplication"
	result := app.SharedConfig()
	if result != expected {
		t.Errorf("expected %q, found %q", expected, result)
	}
}

func TestSharedCache(t *testing.T) {
	expected := "/Library/Caches/TestApplication"
	result := app.SharedCache()
	if result != expected {
		t.Errorf("expected %q, found %q", expected, result)
	}
}

func TestSharedLogs(t *testing.T) {
	expected := "/Library/Logs/TestApplication"
	result := app.SharedLogs()
	if result != expected {
		t.Errorf("expected %q, found %q", expected, result)
	}
}
