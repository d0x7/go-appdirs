# go-appdirs [![GoDoc](https://godoc.org/github.com/emersion/go-appdir?status.svg)](https://godoc.org/xiam.li/appdirs)

Go package to get application directories such as config and cache.

| Platform             | [Linux/BSDs]                             | [macOS]                             | Windows     |
|----------------------|------------------------------------------|-------------------------------------|:------------|
| User-specific data   | `$XDG_DATA_HOME` OR `$HOME/.local/share` | `$HOME/Library/Application Support` | Unsupported |
| User-specific config | `$XDG_CONFIG_HOME` OR `$HOME/.config`    | `$HOME/Library/Application Support` | Unsupported |
| User-specific cache  | `$XDG_CACHE_HOME` OR `$HOME/.cache`      | `$HOME/Library/Caches`              | Unsupported |
| User-specific logs   | `$XDG_CACHE_HOME` OR `$HOME/.cache`      | `$HOME/Library/Logs`                | Unsupported |
| Shared data          | `$XDG_DATA_DIRS` OR `/usr/local/share`   | `/Library/Application Support`      | Unsupported |
| Shared config        | `$XDG_CONFIG_DIRS` OR `/etc/xdg`         | `/Library/Application Support`      | Unsupported |
| Shared cache         | `/tmp/$(APP)-cache`                      | `/Library/Caches`                   | Unsupported |
| Shared logs          | `$(SHARED_DATA)/logs`                    | `/Library/Logs`                     | Unsupported |

[Linux/BSDs]: https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html
[macOS]: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/FileSystemOverview/FileSystemOverview.html#//apple_ref/doc/uid/TP40010672-CH2-SW1

Forked from [emersion](https://github.com/emersion/go-appdir), which got forked from [ProtonMail](https://github.com/ProtonMail/go-appdir)'s implemention, which was inspired by [configdir](https://github.com/shibukawa/configdir).

## Usage

```go
package main

import (
	"os"
	"path/filepath"

	"xiam.li/appdirs"
)

func main() {
	// Get directories for our app
	dirs := appdir.New("my-awesome-app")

	// Get user-specific config dir
	p := dirs.UserConfig()

	// Create our app config dir
	if err := os.MkdirAll(p, 0755); err != nil {
		panic(err)
	}

	// Now we can use it
	f, err := os.Create(filepath.Join(p, "config-file"))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write([]byte("<3"))
}
```

## License

MIT
