local utils = {}

function utils.getText(id)
    return utils.__(id)
end

function utils.getTextF(id, params)
    return utils._f(id, params)
end

function utils.__(id)
    return __KAREN_GETTEXT__(id)
end

function utils._f(id, params)
    return __KAREN_GETTEXT_F__(id, params)
end

return utils
