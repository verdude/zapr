.PHONY: list
list:
	go list -m github.com/verdude/zapr@$(shell git describe --tags --abbrev=0)
