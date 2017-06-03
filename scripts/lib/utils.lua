--
-- Convenience wrappers around the global utility functions
--

local utils = {}

function utils.getText(id)
    return utils.__(id)
end

function utils.getTextF(id, params)
    return utils._f(id, params)
end

function utils.__(id)
    return ____KAREN_UTILS_GETTEXT____(id)
end

function utils._f(id, params)
    return ____KAREN_UTILS_GETTEXT_F____(id, params)
end

return utils
