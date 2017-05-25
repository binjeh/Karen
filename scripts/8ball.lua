local karen = require("karen")
local utils = require("utils")

karen.registerComplex(
    "8ball",
    {"8ball", "8"},
    function(e, user, message)
        if string.len(message) < 3 then
            return utils.__("triggers.8ball.ask_a_question")
        end

        return ":8ball: " + utils.__("triggers.8ball")
    end
)
