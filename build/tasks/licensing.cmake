ADD_CUSTOM_TARGET(license-header
    COMMAND copyright-header
        -w 80
        -a
            src/:
            cmake/:
            assets/:
            scripts/:
            ./*.sh:
            ./CMakeLists.txt:
            ./config.dist.toml:
            docker-compose.yml
        -o .
        -c ./build/c-syntax.yml
        --license-file ./build/AGPL3.erb
)
