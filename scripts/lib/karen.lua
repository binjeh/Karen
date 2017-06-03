local karen = {}

function karen.registerReply(name, listeners, replyId)
    return ____KAREN_API_REGISTER_REPLY____(name, listeners, replyId)
end

function karen.registerComplex(name, listeners, replyFunc)
    return ____KAREN_API_REGISTER_COMPLEX_REPLY____(name, listeners, replyFunc)
end

return karen
