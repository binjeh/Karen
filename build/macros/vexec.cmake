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

# Helper function to save the output of a program in a variable
function(VEXEC)
    # Parse arguments
    cmake_parse_arguments(
        PARSED_ARGS
        ""
        "OUTPUT"
        "COMMAND"
        ${ARGN}
    )

    # Check if required args are present
    if(NOT PARSED_ARGS_OUTPUT)
        message(FATAL_ERROR "[VEXEC] You must provide an output variable!")
    endif()

    if(NOT PARSED_ARGS_COMMAND)
        message(FATAL_ERROR "[VEXEC] You must provide a command list!")
    endif()

    # Allocate temporary result var
    set(mvexec_tmp ?)

    # Check if the program exists on this PC
    list(GET PARSED_ARGS_COMMAND 0 PROGRAM_NAME)
    list(REMOVE_AT PARSED_ARGS_COMMAND 0)
    find_program(PATH__${PROGRAM_NAME} ${PROGRAM_NAME})

    if(NOT PATH__${PROGRAM_NAME})
        message(FATAL_ERROR "[VEXEC] The command ${PROGRAM_NAME} is not installed on this machine!")
    else()
        # Create an alias to the program's path
        set(PROGRAM ${PATH__${PROGRAM_NAME}})

        # If there are no arguments, strip them from the command
        list(LENGTH PARSED_ARGS_COMMAND PARSED_ARGS_COMMAND_LENGTH)

        if(PARSED_ARGS_COMMAND_LENGTH EQUAL 0)
            execute_process(
                COMMAND ${PROGRAM}
                OUTPUT_VARIABLE mvexec_tmp
                ERROR_QUIET
                OUTPUT_STRIP_TRAILING_WHITESPACE
                TIMEOUT 10
            )
            message(STATUS "[KAREN] [VEXEC] [+${PROGRAM_NAME}] ${PARSED_ARGS_OUTPUT} = ${mvexec_tmp}")
        else()
            execute_process(
                COMMAND ${PROGRAM} ${PARSED_ARGS_COMMAND}
                OUTPUT_VARIABLE mvexec_tmp
                ERROR_QUIET
                OUTPUT_STRIP_TRAILING_WHITESPACE
                TIMEOUT 10
            )
            message(STATUS "[KAREN] [VEXEC] [+${PROGRAM_NAME}] ${PARSED_ARGS_OUTPUT} = ${mvexec_tmp}")
        endif()
    endif()

    # Set the return value
    set(${PARSED_ARGS_OUTPUT} ${mvexec_tmp} PARENT_SCOPE)
endfunction()
