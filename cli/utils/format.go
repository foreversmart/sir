package utils

type _Format struct{}

var Format _Format

func (_ *_Format) Enabled(enabled bool) string {
	if enabled {
		return "启用"
	}
	return "禁用"
}
