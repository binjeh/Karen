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
