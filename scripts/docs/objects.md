## Scripting Karen | Objects

Under some circumstances Karen's API takes and/or returns objects to shorten function calls.<br>
Since we're in lua, objects are actually just tables with a predefined structure.

Please note that:
- **ALL** objects are data-only and will never contain any kind of logic, bound functions or channels.
- Cursive members are optional (may be nil at runtime)
- Bold members are mandatory (always have a value, no exceptions)

### User
Represents *any* kind of user.<br>

- **string ID**
- string Email
- **string Username**
- **string Avatar**
- **string Discriminator**
- string Token
- bool Verified
- bool MFAEnabled
- bool Bot
