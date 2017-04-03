container-build:
	docker build --squash -t sn0w/karen-build -f build/build.dockerfile .

container-runtime:
	docker build --squash -t sn0w/karen-runtime -f build/runtime.dockerfile .

containers: container-build container-runtime
