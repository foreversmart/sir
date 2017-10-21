package utils

type _Format struct{}

var Format _Format

func (_ *_Format) Enabled(enabled bool) string {
	if enabled {
		return "启用"
	}
	return "禁用"
}

func (_ *_Format) KVMap(env map[string]string) []string {
	out := []string{}
	for k, v := range env {
		out = append(out, k+": "+v)
	}

	return out
}

func (_ *_Format) KV(key string, value string) string {
	return key + ": " + value
}
