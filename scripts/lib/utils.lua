--
-- Copyright (C) 2015-2017 Lukas Breuer. All rights reserved.
--
-- This file is a part of the Karen Discord-Bot Project ("Karen").
--
-- Karen is free software: you can redistribute it and/or modify
-- it under the terms of the GNU Affero General Public License as published by
-- the Free Software Foundation, either version 3 of the License,
-- or (at your option) any later version.
--
-- Karen is distributed in the hope that it will be useful,
-- but WITHOUT ANY WARRANTY; without even the implied warranty of
-- MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
--
-- See the GNU Affero General Public License for more details.
-- You should have received a copy of the GNU Affero General Public License
-- along with this program. If not, see <http://www.gnu.org/licenses/>.
--

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
