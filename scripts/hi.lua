--[[
--
-- Full credit to Der-Eddy and his original python implementation for Shinobu-Chan.
-- @link https://github.com/Der-Eddy/discord_bot
--
--]]

require("karen").registerReply(
    "git",
    {
        "wave",
        "hi",
        "hello",
        "ohai",
        "ohayou"
    },
    ":wave: " .. require("utils").__("scripts.hi.link")
)
