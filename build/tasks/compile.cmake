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

vexec(OUTPUT KAREN_DYN_VERSION    COMMAND git describe --tags)
vexec(OUTPUT KAREN_DYN_BUILD_TIME COMMAND date "+%T-%D")
vexec(OUTPUT KAREN_DYN_BUILD_USER COMMAND whoami "")
vexec(OUTPUT KAREN_DYN_BUILD_HOST COMMAND hostname "")

ADD_CUSTOM_TARGET(compile-release
    DEPENDS glide-install assets configure
    COMMAND go build -v -o karen
            --ldflags=\"
                -X code.lukas.moe/x/karen/src/version.BOT_VERSION=${KAREN_DYN_VERSION}
                -X code.lukas.moe/x/karen/src/version.BUILD_TIME=${KAREN_DYN_BUILD_TIME}
                -X code.lukas.moe/x/karen/src/version.BUILD_USER=${KAREN_DYN_BUILD_USER}
                -X code.lukas.moe/x/karen/src/version.BUILD_HOST=${KAREN_DYN_BUILD_HOST}
            \"
            ./src
)

ADD_CUSTOM_TARGET(compile-debug
    DEPENDS glide-install assets configure
    COMMAND go build -v -o karen
            --ldflags=\"
                -X code.lukas.moe/x/karen/src/version.BOT_VERSION=${KAREN_DYN_VERSION}
                -X code.lukas.moe/x/karen/src/version.BUILD_TIME=${KAREN_DYN_BUILD_TIME}
                -X code.lukas.moe/x/karen/src/version.BUILD_USER=${KAREN_DYN_BUILD_USER}
                -X code.lukas.moe/x/karen/src/version.BUILD_HOST=${KAREN_DYN_BUILD_HOST}
            \"
            ./src
)

ADD_CUSTOM_TARGET(compile
    DEPENDS compile-release
)

ADD_CUSTOM_TARGET(run
    DEPENDS compile
    COMMAND ./karen
)

ADD_CUSTOM_TARGET(run-debug
    DEPENDS compile-debug
    COMMAND ./karen
)
