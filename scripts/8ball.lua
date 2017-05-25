local utils = require("utils")

require("karen").registerComplex(
    "8ball",
    { "8ball", "8b" },
    function(author, caller, content)
        if string.len(content) < 3 then
            return utils.__("triggers.8ball.ask_a_question")
        end

        return ":8ball: " + utils.__("triggers.8ball")
    end
)
