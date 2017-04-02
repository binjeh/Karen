BOT_VERSION=$(shell git describe --tags)
BUILD_TIME=$(shell date +%T-%D)
BUILD_USER=$(shell whoami)
BUILD_HOST=$(shell hostname)

release: assets_update
	go build -v -o karen \
		--ldflags=" \
			-X code.lukas.moe/x/karen/src/version.BOT_VERSION=$(BOT_VERSION) \
			-X code.lukas.moe/x/karen/src/version.BUILD_TIME=$(BUILD_TIME) \
			-X code.lukas.moe/x/karen/src/version.BUILD_USER=$(BUILD_USER) \
			-X code.lukas.moe/x/karen/src/version.BUILD_HOST=$(BUILD_HOST) \
		" \
		./src

debug: assets_update
	go build -v -race -o karen \
		--ldflags=" \
			-X code.lukas.moe/x/karen/src/version.BOT_VERSION=$(BOT_VERSION) \
			-X code.lukas.moe/x/karen/src/version.BUILD_TIME=$(BUILD_TIME) \
			-X code.lukas.moe/x/karen/src/version.BUILD_USER=$(BUILD_USER) \
			-X code.lukas.moe/x/karen/src/version.BUILD_HOST=$(BUILD_HOST) \
		" \
		./src

compile: release
