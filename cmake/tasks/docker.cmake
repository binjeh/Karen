ADD_CUSTOM_TARGET(container_build
    COMMAND docker build --squash -t sn0w/karen-build -f cmake/build.dockerfile .
)

ADD_CUSTOM_TARGET(container_runtime
    COMMAND docker build --squash -t sn0w/karen-runtime -f cmake/runtime.dockerfile .
)

ADD_CUSTOM_TARGET(containers
    DEPENDS container_build container_runtime
)
