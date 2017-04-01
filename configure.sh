#!/usr/bin/env bash

set -x

# Only continue if ERB is present
command -v erb || {
    echo "Please install ruby and ERB"
    exit 1
}

# Delete old makefile
if [[ -f Makefile ]]; then
    rm Makefile
fi

# Build base makefile
erb build/Makefile.am > Makefile

# Include jobs
for job in build/jobs.d/*.mk; do
    erb ${job} >> Makefile
    echo "" >> Makefile
done

# Add PHONY targets
echo "PHONY:" >> Makefile
for job in build/jobs.d/*.mk; do
    job=${job##*/}
    job=${job%.mk}
    sed -i "/^PHONY:/ s/\$/ $job/" Makefile
done
