#!/usr/bin/env bash

gpp \
    +z \
    -x \
    -U "" "" "(" "," ")" "(" ")" "//#" "\\" \
    -M "//#" "\n" " " " " "\n" "(" ")" \
    --nostdinc \
    --includemarker "/* ___INCLUSION_BOUNDARY___ | Line:% | File:% | % */" \
    "${@:2}" \
    $1
