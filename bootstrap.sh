#!/usr/bin/env bash

#
# Copyright (C) 2015-2017 Lukas Breuer. All rights reserved.
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

set -e

function require {
    command -v $1 2>&1 1>/dev/null || {
        echo "Please install $1 before running this script."
        exit 1
    }
}

set -x

require go
require curl
require glide

go get -v -u golang.org/x/tools/cmd/goimports
go get -v -u github.com/lestrrat/go-bindata/...
go get -v -u git.lukas.moe/sn0w/ropus
