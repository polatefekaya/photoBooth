.DEFAULT_GOAL := run

stage = dev

targetPlatform = windows
targetArch = amd64

targetCombined = $(targetPlatform)-$(targetArch)

targetPath = ./bin/$(stage)/photoBooth
mainFile = cmd/main/main.go

serverFile = ./server/main.py

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

.PHONY: serverDep
serverDep:
	cd ./server
	pip install --upgrade onnxruntime
	pip install --upgrade rembg
	pip install --upgrade Pillow
	pip install --upgrade pika
	cd .

.PHONY: server
server:
	py $(serverFile) 