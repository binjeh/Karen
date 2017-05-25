--[[
--
-- Full credit to Der-Eddy and his original python implementation for Shinobu-Chan.
-- @link https://github.com/Der-Eddy/discord_bot
--
--]]

local utils = require("utils")

require("karen").registerReply(
    "hyperdimension-neptunia",
    {
        "nep",
        "nepgear",
        "neptune"
    },
    utils.__("triggers.nep.text") .. "\n" .. utils.__("triggers.nep.link")
)
