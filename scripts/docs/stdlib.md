## Scripting Karen | StdLib

All the functions listed below are available in the global scope at any time during the execution of your script.

### `string __(id)`
Resolves the given ID to a localized (and maybe randomized) string.<br>
If the ID cannot be resolved the function will return it's input.

### `string _f(id, params)`
Works like `__(id)` but also takes a table of replacements for printf replacements.

Example:
```lua
-- Example translation: "de.hello.world" -> "Hallo %s!"

__f("de.hello.world", {"Bob"})

-- ^ Would return "Hallo Bob!"
```
