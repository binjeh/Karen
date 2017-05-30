ADD_CUSTOM_TARGET(assets-remove
    COMMAND if [ -d src/assets ]\; then rm -r src/assets\; fi
)

ADD_CUSTOM_TARGET(assets-create
    DEPENDS assets-remove
    COMMAND if [ ! -d src/assets ]\; then mkdir src/assets\; fi
    COMMAND go-bindata -pkg assets -o src/assets/assets.go assets/
)

ADD_CUSTOM_TARGET(assets
    DEPENDS assets-create
)
