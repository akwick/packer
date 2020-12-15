package plugin

// +build aix darwin dragonfly freebsd js,wasm linux netbsd openbsd solaris

var (
	FileExtension = ".0_x" + APIVersion // OS-Specific plugin file extention
)
