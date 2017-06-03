ADD_CUSTOM_TARGET(container-build
    COMMAND docker build --squash -t sn0w/karen-build -f cmake/build.dockerfile .
)

ADD_CUSTOM_TARGET(container-runtime
    COMMAND docker build --squash -t sn0w/karen-runtime -f cmake/runtime.dockerfile .
)

ADD_CUSTOM_TARGET(containers
    DEPENDS container-build container-runtime
)
