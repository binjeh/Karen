Origin: https://git.lukas.moe/sn0w/Karen<br>
Mirror: https://github.com/sn0w/Karen

<hr>

<p align="center">
  <a href="https://git.lukas.moe/sn0w/Karen/commits/master">
    <img alt="build status" src="https://git.lukas.moe/sn0w/Karen/badges/master/build.svg" />
  </a>
  <a href="/">
    <img src="https://img.shields.io/github/tag/sn0w/karen.svg?style=flat-square" alt="GitHub tag"/>
  </a>
  <a href="https://goreportcard.com/report/github.com/sn0w/Karen">
    <img src="https://goreportcard.com/badge/github.com/sn0w/Karen?style=flat-square" alt="Go Report Card"/>
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

### Invite
Public invite link is coming after the private beta ends.

Want access?<br>
Register here: https://goo.gl/forms/9J9GYMg8c9IM6a5Z2

### Can I `go get` this bot?
Short answer: No you can't.

Long answer:<br>
No you can't because `go get` sucks.

Huge projects like this need proper version locking.<br>
gopkg.in is a nice try but not good enough.<br>

I use https://glide.sh

### Can I `go build` this bot?
Short answer: No you can't.

Long answer:<br>
No you can't because `go build` sucks.

I use a sophisticated and completely handcrafted build system that uses modular makefiles at it's core.<br>
Thus this is most likely the only go-project that supports c-style macros in comments and conditional compilation.<br>
Are we awesome yet?

### Can I host my own version of this bot?
No.<br>
There will be neither guides nor setup help.<br>
The code is provided for transparency and community-contributions only.
<br>
If you decide to fork Karen remember to comply with the [LICENSE](https://github.com/sn0w/karen/blob/master/LICENSE).

### Why are you saying `high performance` all the time?
I've built a few bots already and all of them were far from performant.<br>
Why do we live in a time where it's acceptable that an EMPTY Java class consumes 10mb RAM?<br>
Why does no one care about optimizing anymore?

That's why I'm building Karen.<br>
I want to create a bot that:

 - Can handle an almost infinite amount of joined guilds
 - Is able to scale vertically **and** horizontally
 - Can play music for free, forever. (Not like that freemium stuff Mee6 does)
 - Will **never** use more resources than absolutely needed
 - Never crashes (or to be precise: be able to recover from almost any `panic()`)

To archieve these goals, the basic rules of coding Karen are:
 - Plugins are compiled into the bot instead of lazy-loading
 - All commands are stateless
 - If a plugin cannot work without states (like `music.go`) it has to implement the state-handling itself
 - Write as much async code as possible
 - Channels > Callbacks
 - Write optimized helper functions instead of duplicated boilerplate code
 - `panic()` on **any** error that is not user-related and `recover()` later
 - Mind your CPU! A coroutine should die as soon as possible
 - If you wait for something in an endless loop let it `sleep()` as long as possible between iterations.
 - A function's cyclomatic complexity should be as low as possible
 - A function should have as few responsibilities as possible

### Achievements

- Never exceeded 2% CPU usage at the time of writing.
- Never used more than 8MB of it's allocated heap.

![](https://i.imgur.com/lGf08Yo.png)

### Docs
Hancrafted guide soon (tm)

Until then use GoWalker/GoDoc for coding guides and
the homepage for usage help.

### Disclaimer
This bot is still in a early stage.<br>
Please expect (rare) crashes and minor performance problems until the bot is mature enough.

### About the bad GoReportCard score
This project will probably never reach more than a D on GoReportCard because I hugely disagree on gofmt's decision to use tabs instead of spaces. I know that this discussion is as old as programming itself, but this is my opinion on this. I prefer spaces. Always. I might create a fork of gofmt that allows spaces at some point.

The primary reason for having GoReportCard at all are the go_vet, gocyclo, ineffassign and misspell tests.<br>
You are forced to keep the score as-is or improve it when making contributions to the project.

### How does this work?
I've drawn a colorful picture just for you!

![](http://i.imgur.com/lI3VJDo.png)
