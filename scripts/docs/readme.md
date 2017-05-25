## Scripting Karen

You're able to script Karen by using lua - a small nifty programming language that we embedded into the bot's core. This functionality was added because users started to request plugins that do *extremely* simple things like sending a link or a random image. Coding these plugins required a lot of boilerplate code.

Scripting simplifies this, by hiding the duplicated code behind lua's VM.<br>
It also opens the possibility to develop functionality, for a very large and active community of lua scripters.

Every lua file in the `_scripts` folder is interpreted at boot and then attached to a golang event listener.

### Further Docs:
- #### [API Docs](api.md)
- #### [Passed Objects](objects.md)
- #### [StdLib](stdlib.md)
