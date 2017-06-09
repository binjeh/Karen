<p align="center">
  <img alt="karen" width="96" src="http://i.imgur.com/VfgrwQz.jpg">
  <br>
  <a href="<https://travis-ci.org/SubliminalHQ/karen">
    <img alt="build status" src="https://travis-ci.org/SubliminalHQ/karen.svg?branch=master" />
  </a>
  <a href="https://www.codacy.com/app/lukas-breuer/karen?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=sn0w/karen&amp;utm_campaign=Badge_Grade">
    <img src="https://api.codacy.com/project/badge/Grade/ec90cbf66f5c4ecbab15d5dfe73c0ddd"/>
  </a>
  <a href="https://codebeat.co/projects/github-com-sn0w-karen-master">
    <img alt="codebeat badge" src="https://codebeat.co/badges/2d89b948-da1b-40e4-8ecd-f3b7dc394591" />
  </a>
  <a href="#">
    <img src="https://img.shields.io/github/tag/sn0w/karen.svg?style=flat-square" alt="GitHub tag"/>
  </a>
  <a href="https://gowalker.org/github.com/sn0w/Karen">
    <img src="http://gowalker.org/api/v1/badge" alt="Go Walker" />
  </a>
  <br>
  Karen is a highly efficient, multipurpose Discord bot written in Golang.
  <br>
  <br>
  Got any problems or just want to chat with me and other devs?<br>
  Join the Discord Server! :)<br>
  <a href="https://discord.karen.vc">
    <img src="https://discordapp.com/api/guilds/180818466847064065/widget.png">
  </a>
</p>
<hr/>

### Docs
Hancrafted guide soon (tm)

Until then use GoWalker/GoDoc for coding guides and
the homepage for usage help.

### Disclaimer
This bot is still in a early stage.<br>
Please expect (rare) crashes and minor performance problems until the bot is mature enough.

### Karen's Friends :tada:

Bots built by friends or awesome strangers

|Avatar|Name|Language|Link|
|:-:|:-:|:-:|:-:|
|![](http://i.imgur.com/SrgZI3g.png)|Emily|Java|[MaikWezinkhof/DiscordBot](https://github.com/MaikWezinkhof/DiscordBot)
|![](https://cdn.discordapp.com/avatars/270931284489011202/b7b1f9820c4751ffa3d0e11c97bc2f38.png?size=64)|Sora|C#|[Serenity/Sora](http://git.argus.moe/serenity/SoraBot)
|![](https://cdn.discordapp.com/avatars/260867076774821899/2dda452db1e35f833a187df9dd4f1749.png?size=64)|Nep|C#|[Serraniel/Nep-Bot](https://github.com/Serraniel/Nep-Bot)
|![](http://i.imgur.com/Tb0FZoZ.png)|Shinobu-Chan|Python 3|[Der-Eddy/discord_bot](https://github.com/Der-Eddy/discord_bot)

### Can I `go get` this bot?
Short answer: No you can't. `go get` sucks.

Huge projects like this need proper version locking.<br>
gopkg.in is a nice try but not good enough.<br>

We use https://glide.sh

### Can I `go build` this bot?
Short answer: No you can't. `go build` sucks.

We use CMake, GNU Make and some shell magic instead.
Thus this is most likely the only golang project on earth that supports c-style macros in comments and conditional compilation.<br>

Are we awesome yet?

### Can you help me self-hosting Karen?
No.<br>
You are allowed to host Karen and fork the project (given that you comply to the license),
but there will be neither guides nor setup help.<br>

### Why are you saying `high performance` all the time?
Why do we live in a time where it's acceptable that an EMPTY Java class consumes 10mb RAM?<br>
Why does no one care about optimizing anymore?

Our vision is to create a bot that:

 - Can handle an almost infinite amount of joined guilds (aka good sharding)
 - Is able to scale vertically and horizontally
 - Can play music for free, forever. (Not like that freemium stuff Mee6 does)
 - Will never use more resources than absolutely needed
 - Never crashes

### Achievements

- Never exceeded 2% CPU usage at the time of writing.
- Never used more than 8MB of it's allocated heap.

![](https://i.imgur.com/lGf08Yo.png)
