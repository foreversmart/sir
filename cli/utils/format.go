package utils

type _Format struct{}

var Format _Format

func (_ *_Format) Enabled(enabled bool) string {
	if enabled {
		return Style.Success("enabled")
	}
	return Style.Disabled("disabled")
}

func (_ *_Format) KV(key string, value string) string {
	return key + ": " + value
}
