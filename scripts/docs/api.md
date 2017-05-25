## Scripting-Karen | API

There are many ways to register your plugin depending on it's purpose and complexity. 
You have to choose yourself which one of them fits best.

### `bool RegisterReply(name, listeners, replyId)`
This is the most basic type of script you can create. It takes a name (string), a table of listeners and a string that will be localized through `GetText()` and the replied.

Example:
```lua
-- Registers a script named "ping-pong" that sends the
-- text behind scripts.pingpong.reply when [p]ping or [p]
-- appear in the chat.
RegisterReply(
    "ping-pong",
    {"ping", "p"},
    "scripts.pingpong.reply"
)
```

### `bool RegisterComplexReply(name, listeners, replyFunc)`

Works exactly like `RegisterReply()` but takes a function instead of a simple string.

Parameters of  `replyFunc`:
- `string` The listener that triggered this execution
- `User` The message's author
