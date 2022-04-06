package main

import "embed"

//go:embed html/*
var webUIFiles embed.FS

func WebUI(path string) ([]byte, bool) {
	ret, err := webUIFiles.ReadFile(path)
	if err != nil {
		return nil, false
	}
	return ret, true
}
