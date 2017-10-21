package config

var HOST = "http://localhost:8080"

func ApiPath(subPath string) string {
	return HOST + subPath
}
