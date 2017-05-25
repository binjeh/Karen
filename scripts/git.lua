require("karen").registerReply(
    "git",
    {
        "git",
        "gitlab",
        "github",
        "repo"
    },
    ":earth_africa: " .. require("utils").__("triggers.git")
)
