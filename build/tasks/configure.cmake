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

set(EXCLUDE_MUSIC 0 CACHE STRING "Exclude the musicbot? (0/1)")
set(EXCLUDE_RADIO 0 CACHE STRING "Exclude radio streaming? (0/1)")
set(EXCLUDE_PLUGINS 0 CACHE STRING "Exclude plugins? (0/1)")
set(EXCLUDE_SCRIPTING 0 CACHE STRING "Exclude scripting engine? (0/1)")

ADD_CUSTOM_TARGET(configure
    COMMAND ./ppw.sh \"
        -DEXCLUDE_MUSIC=${EXCLUDE_MUSIC}
        -DEXCLUDE_RADIO=${EXCLUDE_RADIO}
        -DEXCLUDE_PLUGINS=${EXCLUDE_PLUGINS}
        -DEXCLUDE_SCRIPTING=${EXCLUDE_SCRIPTING}
    \"
)
