BOT_VERSION=$(shell git describe --tags)
BUILD_TIME=$(shell date +%T-%D)
BUILD_USER=$(shell whoami)
BUILD_HOST=$(shell hostname)

release: assets_update
	go build -v -o karen \
		--ldflags=" \
			-X git.lukas.moe/sn0w/Karen/x/version.BOT_VERSION=$(BOT_VERSION) \
			-X git.lukas.moe/sn0w/Karen/x/version.BUILD_TIME=$(BUILD_TIME) \
			-X git.lukas.moe/sn0w/Karen/x/version.BUILD_USER=$(BUILD_USER) \
			-X git.lukas.moe/sn0w/Karen/x/version.BUILD_HOST=$(BUILD_HOST) \
		" \
		./x

debug: assets_update
	go build -v -race -o karen \
		--ldflags=" \
			-X git.lukas.moe/sn0w/Karen/x/version.BOT_VERSION=$(BOT_VERSION) \
			-X git.lukas.moe/sn0w/Karen/x/version.BUILD_TIME=$(BUILD_TIME) \
			-X git.lukas.moe/sn0w/Karen/x/version.BUILD_USER=$(BUILD_USER) \
			-X git.lukas.moe/sn0w/Karen/x/version.BUILD_HOST=$(BUILD_HOST) \
		" \
		./x

compile: release
