--
-- Convenience wrappers around the global register-api
--

local karen = {}

function karen.registerReply(name, listeners, replyId)
    return __KAREN_REGISTER_REPLY__(name, listeners, replyId)
end

function karen.registerComplex(name, listeners, replyFunc)
    return __KAREN_REGISTER_COMPLEX__(name, listeners, replyFunc)
end

return karen
