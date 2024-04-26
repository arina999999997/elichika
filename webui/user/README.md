# User features for WebUI
## Authentication
The following requirements are needed for authentication:

- User need to login via user id and password (empty string if they haven't set one)
- User must have completed the tutorial

The requirements are necessary for some feature, but they also present spamming (but maybe not a lot).

## Locale
When using the WebUI, it's important to choose the correct client language. This is mainly for the server, as some content might not exist for other languages and will make your account unavailable.

In the future, maybe the client language can also change how the WebUI display things.

## Features
Some features and what they do. Viewing them on the webui also explain the feature a bit more.
### Account builder
This allow you to do things fast in the game, it also has a button to get you a maxed out account (based on your locale).
### Resource helper
This allow you to get items in the game so you can progress faster, but it won't do the thing for you like the account builder.
### Reset progress
This allow you to reset and play stuff again, while keeping the other data of your account.
### Import/Export account
Allow you to export / import your account:

- Exporting to .db file allow you to save all data that is compatible with elichika.
- Exporting to .json allow you to get the canonical network data, so you can import them elsewhere if necessary.

Importing an account will overwrite your (account logged into webui's) current progress.

### Advanced
Some features require advancded understanding of the game's internal (or the server too).

- Add item using id: Allow user to add an item directly using id.