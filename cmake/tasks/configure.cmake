set(EXCLUDE_MUSIC 0 CACHE BOOL "Compile the musicbot?")
set(EXCLUDE_RADIO 0 CACHE BOOL "Compile radio streaming?")
set(EXCLUDE_PLUGINS 0 CACHE BOOL "Compile plugins?")
set(EXCLUDE_SCRIPTING 0 CACHE BOOL "Compile scripting engine?")

ADD_CUSTOM_TARGET(configure
    COMMAND ./ppw.sh \"
        -DEXCLUDE_MUSIC=${EXCLUDE_MUSIC}
        -DEXCLUDE_RADIO=${EXCLUDE_RADIO}
        -DEXCLUDE_PLUGINS=${EXCLUDE_PLUGINS}
        -DEXCLUDE_SCRIPTING=${EXCLUDE_SCRIPTING}
    \"
)
