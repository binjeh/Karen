.PHONY: container-build
container-build:
	docker build --squash -t sn0w/karen-build -f build/build.dockerfile .

.PHONY: container-runtime
container-runtime:
	docker build --squash -t sn0w/karen-runtime -f build/runtime.dockerfile .

.PHONY: containers
containers: container-build container-runtime
