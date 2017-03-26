# How to contribute

Third-party patches are essential for keeping Open-Source software awesome.
We simply can't think of all circumstances and myriad configurations for
building and running our software. We want to keep it as easy as possible to
contribute changes that get things working in your environment or add new awesome
features. There are a few guidelines that we need contributors to follow so that
we can have a chance of keeping on top of things.

## Getting Started

* Make sure you have a registered account at Lukas' Git (https://git.lukas.moe).
* Submit an issue that describes you problem or idea, assuming one does not already exist.
  * Clearly describe the issue/idea including steps to reproduce when it is a bug.
  * Make sure you fill in the earliest version that you know has the issue if you're reporting a bug.
* ***WAIT*** until someone tells you that we approve your feature request or bug report.

## Making Changes

* Fork the repository
* Create a topic branch from where you want to base your work.
  * This is usually the master branch.
  * Only target release/production branches if you are certain your fix must be on that branch and talked about that with us.
  * To quickly create a topic branch based on master do a `git checkout -b feature/my_awesome_feature`.
  * Please avoid working directly on the `master` branch. Your request might be rejected otherwise.
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
