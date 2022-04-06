package src

var loadWebUI func(string) ([]byte, bool)

func webUI(path string) ([]byte, bool) {
	return loadWebUI(path)
}

func SetWebUI(loadWebUIFunc func(path string) ([]byte, bool)) {
	loadWebUI = loadWebUIFunc
}
