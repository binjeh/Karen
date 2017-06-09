#
# Karen - A highly efficient, multipurpose Discord bot written in Golang
#
# Copyright (C) 2015-2017 Lukas Breuer
# Copyright (C) 2017 Subliminal Apps
#
# This file is a part of the Karen Discord-Bot Project ("Karen").
#
# Karen is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License,
# or (at your option) any later version.
#
# Karen is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
#
# See the GNU Affero General Public License for more details.
# You should have received a copy of the GNU Affero General Public License
# along with this program. If not, see <http://www.gnu.org/licenses/>.
#

include(build/macros/vexec.cmake)

# Create variables that will be compiled into the bot later
vexec(OUTPUT KAREN_DYN_VERSION    COMMAND git describe --tags)
vexec(OUTPUT KAREN_DYN_BUILD_TIME COMMAND date "+%T-%D")
vexec(OUTPUT KAREN_DYN_BUILD_USER COMMAND whoami "")
vexec(OUTPUT KAREN_DYN_BUILD_HOST COMMAND hostname "")

# Helper function to create a go compilation target
function(ADD_COMPILER_TASK)
    # Parse arguments
    cmake_parse_arguments(
        PARSED_ARGS
        ""
        "NAME;TARGET"
        "DEPENDS;FLAGS;ALIASES"
        ${ARGN}
    )

    # Check if required args are present
    if(NOT PARSED_ARGS_NAME OR NOT PARSED_ARGS_TARGET OR NOT PARSED_ARGS_DEPENDS OR NOT PARSED_ARGS_ALIASES)
        message(FATAL_ERROR "Call to ADD_COMPILER_TASK had incomplete arguments!")
    endif()

    # Log the current invocation
    message(STATUS "[KAREN] [COMPILER] [+] NAME='${PARSED_ARGS_NAME}' ALIASES='${PARSED_ARGS_ALIASES}'")

    # Conditionally pass flags because CMAKE does not like empty flag arguments
    if(NOT PARSED_ARGS_FLAGS)
        ADD_CUSTOM_TARGET(${PARSED_ARGS_NAME}
            DEPENDS ${PARSED_ARGS_DEPENDS}
            COMMAND go build
                    -o ${PARSED_ARGS_TARGET}
                    --ldflags=\"
                    -X code.lukas.moe/x/karen/src/version.BOT_VERSION=${KAREN_DYN_VERSION}
                    -X code.lukas.moe/x/karen/src/version.BUILD_TIME=${KAREN_DYN_BUILD_TIME}
                    -X code.lukas.moe/x/karen/src/version.BUILD_USER=${KAREN_DYN_BUILD_USER}
                    -X code.lukas.moe/x/karen/src/version.BUILD_HOST=${KAREN_DYN_BUILD_HOST}
                    \"
                    ./src
        )
    else()
        ADD_CUSTOM_TARGET(${PARSED_ARGS_NAME}
            DEPENDS ${PARSED_ARGS_DEPENDS}
            COMMAND go build
                    ${PARSED_ARGS_FLAGS}
                    -o ${PARSED_ARGS_TARGET}
                    --ldflags=\"
                    -X code.lukas.moe/x/karen/src/version.BOT_VERSION=${KAREN_DYN_VERSION}
                    -X code.lukas.moe/x/karen/src/version.BUILD_TIME=${KAREN_DYN_BUILD_TIME}
                    -X code.lukas.moe/x/karen/src/version.BUILD_USER=${KAREN_DYN_BUILD_USER}
                    -X code.lukas.moe/x/karen/src/version.BUILD_HOST=${KAREN_DYN_BUILD_HOST}
                    \"
                    ./src
        )
    endif()

    # Register aliases of this task (if needed) by creating new tasks that depend on this task
    foreach(ALIAS ${PARSED_ARGS_ALIASES})
        ADD_CUSTOM_TARGET(${ALIAS}
            DEPENDS ${PARSED_ARGS_NAME}
            COMMAND echo \">>> The executed task '${ALIAS}' was actually an alias for '${PARSED_ARGS_NAME}' <<<\"
        )
    endforeach()
endfunction()
