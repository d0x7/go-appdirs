package appdirs

type dirs struct {
	name string
}

func (d dirs) homeDir() string {
	panic("Windows is not yet implemented")
}

func (d dirs) UserData() string {
	panic("Windows is not yet implemented")
}

func (d dirs) UserConfig() string {
	panic("Windows is not yet implemented")
}

func (d dirs) UserCache() string {
	panic("Windows is not yet implemented")
}

func (d dirs) UserLogs() string {
	panic("Windows is not yet implemented")
}

func (d dirs) SharedData() string {
	panic("Windows is not yet implemented")
}

func (d dirs) SharedConfig() string {
	panic("Windows is not yet implemented")
}

func (d dirs) SharedCache() string {
	panic("Windows is not yet implemented")
}

func (d dirs) SharedLogs() string {
	panic("Windows is not yet implemented")
}
