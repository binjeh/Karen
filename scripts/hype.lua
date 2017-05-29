--[[
--
-- Full credit to Der-Eddy and his original python implementation for Shinobu-Chan.
-- @link https://github.com/Der-Eddy/discord_bot
--
--]]

local utils = require("utils")

require("karen").registerReply(
    "hype",
    { "hype", "hypu" },
    utils.__("scripts.hypetrain.text") .. "\n" .. utils.__("scripts.hypetrain.link")
)
