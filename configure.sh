#!/usr/bin/env bash

### Automake-like checks ###

function println {
    printf "$1\n"
}

function checkPre {
    printf "Checking $1 ... "
}

function checkIf {
    checkPre "if $1"
}

function checkBin {
    checkPre "for $1"
    command -v $1 2>&1 1>/dev/null || {
        echo "Please install '$1' and try again."
        exit 1
    }
    println "found"
}

checkBin go
checkIf "go is version 1.8.*"
if [[ "$(go version)" != *"go1.8"* ]]; then
    println "no"
    exit 1
fi
println "yes"

checkBin ruby
checkBin erb
checkBin gcc
checkBin cpp
checkBin glide
checkBin go-bindata
checkBin ropus
checkBin ffmpeg
checkBin ffprobe

### Dynamic compilation ###
PP_DEFS=""

while test $# -gt 0; do
    case "$1" in
        --without-music)
        PP_DEFS="$PP_DEFS -DEXCLUDE_MUSIC"
        shift
        ;;

        --without-radio)
        PP_DEFS="$PP_DEFS -DEXCLUDE_RADIO"
        shift
        ;;

        --without-modules)
        PP_DEFS="$PP_DEFS -DEXCLUDE_MODULES"
        shift
        ;;

        --without-triggers)
        PP_DEFS="$PP_DEFS -DEXCLUDE_TRIGGERS"
        shift
        ;;

        *)
        break
        ;;
    esac
done

echo
echo "----- Build Flavour -----"
if [[ "$PP_DEFS" == "" ]];then
    echo "DEFAULT"
else
    echo "${PP_DEFS}"
fi
echo "-------------------------"
echo

[[ -f ./src/modules/modules.go ]] && rm ./src/modules/modules.go
./pp.sh ./src/modules/_modules.go ${PP_DEFS} > ./src/modules/modules.go

### Build makefile ###

printf "Building makefile... "

# Delete old makefile
if [[ -f Makefile ]]; then
    rm Makefile
fi

# Build base makefile
cp build/Makefile.mk Makefile

# Include jobs
for job in build/jobs.d/*.mk; do
    erb ${job} >> Makefile
    echo "" >> Makefile
done

# Add PHONY targets
echo ".PHONY:" >> Makefile
for job in build/jobs.d/*.mk; do
    job=${job##*/}
    job=${job%.mk}
    sed -i "/^\.PHONY:/ s/\$/ $job/" Makefile
done

printf "done\n"
