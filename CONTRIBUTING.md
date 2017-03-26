# How to contribute

Third-party patches are essential for keeping Open-Source software awesome.
We simply can't think of all circumstances and myriad configurations for
building and running our software. We want to keep it as easy as possible to
contribute changes that get things working in your environment or add new awesome
features. There are a few guidelines that we need contributors to follow so that
we can have a chance of keeping on top of things.

### Development

We're an open and flexible team of developers and don't force anyone into a strict workflow.
Just make sure that your editor or IDE has a proper syntax-highlighter and linter.
Pull-Requests with syntax errors will be rejected.
**Also PLEASE check that your editor doesn't `gofmt` the code on save**.
We use their coding-style but have 4 spaces instead of tabs.

Recommended IDE's are:
- Jetbrains Gogland
- IntelliJ IDEA
- VS-Code
- Sublime Text 3

We can't help you with this topic though.<br>
***Please don't start working on features before you figured out a workflow that works for you.***

### Requirements

- Hardware
    - A 64bit CPU
    - CPU-Support for AVX2, AVX or SSE3+ instructions is a plus but not (yet) mandatory
- OS
    - Any recent Linux Distro, Windows 7+ or macOS 10.9+
- Shell
    - BASH 4.0 or newer
    - Alternatively any BASH 4.0 compatible shell (like ZSH)
- Software
    -  [git](http://git-scm.com) (Version 2.0 or newer)
    -  [Go](http://golang.org) (Version 1.7 or newer)
- Mandatory Tools
    - [go-bindata](https://github.com/jteeuwen/go-bindata)
    - [Glide](https://glide.sh/)

<br>

If you're a Windows user you need to setup a UNIX-like environment.<br>
That means installing a software like [MINGW](http://www.mingw.org/) or [Cygwin](https://www.cygwin.com/)
and using BASH instead of CMD or PowerShell.

<br>

***PLEASE NOTE:***<br>
Do **NOT** use the "classic" POSIX Shell (`sh`).<br>
Our scripts use advanced parameter-expansion and the double-braced `if` condition (`[[ <condition> ]];`) <br>
and those features are not supported below bash.

## Getting Started

* Make sure you have a registered account at Lukas' Git (https://git.lukas.moe).
* Submit an issue that describes you problem or idea, assuming one does not already exist.
  * Clearly describe the issue/idea including steps to reproduce when it is a bug.
  * Make sure you fill in the earliest version that you know has the issue if you're reporting a bug.
* ***WAIT*** until someone tells you that we approve your feature request or bug report.

## Getting the code

* Make sure that the folder `$GOPATH/src/git.lukas.moe/sn0w` exists
* Fork the repository
* Clone the repository to `$GOPATH/src/git.lukas.moe/sn0w/Karen`
    * This is needed to avoid namespacing issues with Go's compiler.
* Create a topic branch from where you want to base your work.
  * This is usually the master branch.
  * Only target release/production branches if you are certain your fix must be on that branch and talked about that with us.
  * To quickly create a topic branch based on master do a `git checkout -b feature/my_awesome_feature`.
  * Please avoid working directly on the `master` branch. Your request might be rejected otherwise.

## Installing dependencies

The depdencies are currently managed with [Glide](https://glide.sh/).
It's a tool that works just like npm, cargo, composer and all the other fancy dependency managers - but for go.
It's actually a wrapper around Go's vendoring feature but adds critical features like version locking.

After installing it just `cd` into Karen's folder and run `glide install`.

## Making changes
* Make commits of logical units to make rebasing/merging/reviewing easier for us.
* Check for unnecessary whitespace changes with `git diff --check` before committing.
* Make sure your commit messages are in the proper format and not too long.
* Your commit message must not exceed 120 chars.
    * Rule of thumb: Short title, long explanation in body.
* Make sure you have added the necessary tests for your changes.
* Run _all_ the tests to assure nothing else was accidentally broken.
* Build Karen and test the commands on Karen's Sandbox or private guilds.

## Submitting Changes

* Push your changes to a topic branch in your fork.
* Submit a merge request.
* The core team looks at Merge Requests on a regular basis.
Please don't spam comments like `Please merge this already!!!`.
We will merge it when we think it's time to do so.
Excessive spamming increases the chances that we reject your request.
You will get an email when your request was accepted.
* Feedback will be given if necessary
That means we will annotate your code if we see bugs or if you don't use our
style-guidelines properly.
* After feedback has been given we expect responses within two weeks. After two
  weeks we may close the merge request if it isn't showing any activity.
    * The interval can be expanded if you tell us what's going on. Everyone here has a real-life and understands that
      "having time" is pretty uncommon when having a job and/or family.

## Attribution
This document is adapted from [puppetlabs/puppet](https://github.com/puppetlabs/puppet)
