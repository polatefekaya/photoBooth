.DEFAULT_GOAL := run

stage = dev

targetPlatform = windows
targetArch = amd64

targetCombined = $(targetPlatform)-$(targetArch)

targetPath = ./bin/$(stage)/photoBooth
mainFile = cmd/main/main.go

.PHONY: compile
compile:
	echo "compiling for different platforms"
	GOOS=linux GOARCH=$(targetArch) go build -o $(targetPath)-linux-$(targetArch) $(mainFile)
	GOOS=windows GOARCH=$(targetArch) go build -o $(targetPath)-windows-$(targetArch) $(mainFile)

.PHONY: run
run: compile
	echo "Running compiled binary for $(targetCombined)"
	$(targetPath)-$(targetCombined)

.PHONY: clean
clean:
	go clean
	rm $(targetPath)-linux-$(targetArch)
	rm $(targetPath)-windows-$(targetArch)

.PHONY: cleanRun
cleanRun: clean compile run

