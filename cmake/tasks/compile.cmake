# Define dynamically-compiled variables
set(KAREN_DYN_VERSION ?)
set(KAREN_DYN_BUILD_TIME ?)
set(KAREN_DYN_BUILD_USER ?)
set(KAREN_DYN_BUILD_HOST ?)

find_program(GIT git)
if(GIT)
    execute_process(COMMAND git describe --tags OUTPUT_VARIABLE KAREN_DYN_VERSION ERROR_QUIET OUTPUT_STRIP_TRAILING_WHITESPACE)
endif()

find_program(DATE date)
if(DATE)
    execute_process(COMMAND date +%T-%D OUTPUT_VARIABLE KAREN_DYN_BUILD_TIME ERROR_QUIET OUTPUT_STRIP_TRAILING_WHITESPACE)
endif()

find_program(WHO whoami)
if(WHO)
    execute_process(COMMAND whoami OUTPUT_VARIABLE KAREN_DYN_BUILD_USER ERROR_QUIET OUTPUT_STRIP_TRAILING_WHITESPACE)
endif()

find_program(HOST hostname)
if(HOST)
    execute_process(COMMAND hostname OUTPUT_VARIABLE KAREN_DYN_BUILD_HOST ERROR_QUIET OUTPUT_STRIP_TRAILING_WHITESPACE)
endif()

ADD_CUSTOM_TARGET(glide_install
    COMMAND test -d vendor || glide install
)

ADD_CUSTOM_TARGET(compile_release
    DEPENDS glide_install assets configure
    COMMAND go build -v -o karen
            --ldflags=\"
                -X code.lukas.moe/x/karen/src/version.BOT_VERSION='${KAREN_DYN_VERSION}'
                -X code.lukas.moe/x/karen/src/version.BUILD_TIME='${KAREN_DYN_BUILD_TIME}'
                -X code.lukas.moe/x/karen/src/version.BUILD_USER='${KAREN_DYN_BUILD_USER}'
                -X code.lukas.moe/x/karen/src/version.BUILD_HOST='${KAREN_DYN_BUILD_HOST}'
            \"
            ./src
)

ADD_CUSTOM_TARGET(compile_debug
    DEPENDS glide_install assets configure
    COMMAND go build -v -o karen
            --ldflags=\"
                -X code.lukas.moe/x/karen/src/version.BOT_VERSION='${KAREN_DYN_VERSION}'
                -X code.lukas.moe/x/karen/src/version.BUILD_TIME='${KAREN_DYN_BUILD_TIME}'
                -X code.lukas.moe/x/karen/src/version.BUILD_USER='${KAREN_DYN_BUILD_USER}'
                -X code.lukas.moe/x/karen/src/version.BUILD_HOST='${KAREN_DYN_BUILD_HOST}'
            \"
            ./src
)

ADD_CUSTOM_TARGET(compile
    DEPENDS compile_release
)
