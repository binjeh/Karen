## Scripting-Karen | API

There are multiple ways to register your plugin depending on it's purpose and complexity. 
You have to choose yourself which one of them fits best.

### `void RegisterReply(name, listeners, replyId)`
This is the most basic type of script you can create. 
It takes a name (string), a table of listeners and a string that the bot should reply.

Example:
```lua
-- Registers a script named "ping-pong" that sends the
-- text behind scripts.pingpong.reply when [p]ping or [p]
-- appear in the chat.
require("karen").registerReply(
    "ping-pong",
    {"ping", "p"},
    require("utils").GetText("scripts.pingpong.reply")
)
```

### `bool RegisterComplex(name, listeners, replyFunc)`

Works exactly like `RegisterReply()` but takes a function instead of a simple string.

Example:
```lua
-- Store reference to utils
local utils = require("utils")

-- Registers a complex plugin that uses a callback
-- instead of strings. Can be used to implement simple logic
-- like checking the arguments or user permissions.
require("karen").registerComplex(
    "8ball",
    { "8ball", "8b" },
    function(author, caller, content)
        if string.len(content) < 3 then
            return utils.__("triggers.8ball.ask_a_question")
        end

        return ":8ball: " .. utils.__("triggers.8ball")
    end
)

```
