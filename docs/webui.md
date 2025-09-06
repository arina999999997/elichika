
# WebUI
The WebUI allow you to interact with the server in a more direct way, both to change the server and to change your player data.

To use the WebUI, navigate to the relevant address using a web browser.

## User WebUI

The User WebUI's address is: `<server_address>/webui/user`, which default to http://127.0.0.1:8080/webui/user:

- This is the case for the embedded server. You have to keep the client open and open the browser to access this.
- For other setup, it's up to you to figure out what the address is.

### Authentication
You need to login to the WebUI to use it:

- Enter your user id (viewable on the title screen / profile menu) and password (empty string if you haven't set one).

Furthermore, you must have completed the tutorial in the game (skipping is fine). The requirements are necessary for some feature, but they also prevent spamming for multiplayer servers.

### Locale
When using the WebUI, it's important to choose the correct client language. This is mainly for the server, as some content might not exist for other languages and will brick your account.

In the future, maybe the client language can also change how the WebUI display things.

### Features
Some features and what they do. Viewing them on the webui also explain the feature a bit more.
#### Account builder
This allow you to do things fast in the game, it also has a button to get you a maxed out account (based on your locale).
#### Resource helper
This allow you to get items in the game so you can progress faster, but it won't do the thing for you like the account builder.
#### Reset progress
This allow you to reset and play stuff again, while keeping the other data of your account.
#### Import/Export account
Allow you to export / import your account:

- Exporting to .db file allow you to save all data that is compatible with elichika.
- Exporting to .json allow you to get the canonical network data, so you can import them elsewhere if necessary.

Importing an account will overwrite your (account logged into webui's) current progress.

For more details, check [import export docs](https://github.com/arina999999997/elichika/blob/master/docs/import_export.md).

#### Other features
*This is considered [advanced usage](https://github.com/arina999999997/elichika/blob/master/docs/advanced_usage.md)*.

You have access to some other features from the user WebUI, but they required advanced understanding of the game's internal. Be prepared to figure it out with minimal help, otherwise don't touch these features as it might brick your account.

## Admin WebUI

The Admin WebUI's address is: `<server_address>/webui/admin`, which default to http://127.0.0.1:8080/webui/admin:
- This is the case for the embedded server. You have to keep the client open and open the browser to access this.
- For other setup, it's up to you to figure out what the address is.

The admin WebUI is used to change the server's behaviors, to use the admin WebUI, you have to have the admin password:

- By default this is empty (no need to enter anything, just press the login button).
- If you use the embedded version or host your own server for yourself, changing the admin password is not recommended.
- If for some reason you want to change the password, you can change it in the config editor or by editting the config file directly (figure out how yourself).

### Config editor
Detailed explanations of some config options:

- Default item count:
    - The amount of items to give a player to start with.
    - Default to a generous amount, but you can set it to 0 to have a more true experience.
    - Note that you have to obtain the item in game first before you are given the "default item count" amount of that item.
- Mission progress multiplier:
    - How much progress you get for doing 1 thing in the mission.
    - If you want to complete mission faster, set this number to something big.
    - But don't set it too big otherwise you might have some issue.
- Resource config:
    - The config of how the resources work on the server.
    - `original` means the resources behave like it would in the original server. Every action that cost resources will consume those resources.
    - `comfortable` is the default settings. Things like star gems, LP, AP are unlimited. The daily song play limit or the daily tap bond limit is also removed. This is the recommended settings if you just want to play and experience the game without all the money-making limitation.
    - `free` is the free settings. Generally, resources can only go up and not down.
    - Keep in mind that some resources/systems are not controlled by these settings, but they are pretty minor.
    - And this doesn't apply to the accessories (but apply to the accessory items).
- Event frequency:
    - How fast the events cycle around.

Modifying any other field is considered *[advanced usage](https://github.com/arina999999997/elichika/blob/master/docs/advanced_usage.md)*. If you know what you are doing, here's some more explanation for some other field:

- Server's address:

    - The address to host the server at.
    - Default to ``0.0.0.0:8080`` (server listen at port 8080 on all interfaces).

- CDN server's address:

    - The server for clients to download assets.
    - Default to  https://llsifas.imsofucking.gay/static/ (special thanks to sarah for hosting it).
    - `elichika` also has the ability to host the CDN itself:

        - To do this, put the relevant files in `elichika/static`:
          - For each package, the server expect it in `elichika/static/<package_name>`
        - Then set the CDN server address to the STRING (no protocol) `elichika` (or `elichika_tls` if you're using HTTPS).
        - This will automatically use whatever the address the client reach `elichika` with as the CDN server.
        - Aside from that, you can also just use the address like normal.
        - You should look into this if you want to further develop the game/server, as doing so might require redownloading things a lot.
    - You can also use other CDNs, but keep in mind that there are some requirements that need to be met, otherwise some download will result in errors.

### Event Selector

Select the event you want to play. Note that each events need to be recreated, so only the events ready to play are shown.

### Event Scheduler

Schedule the next event you want to play, it will be loaded once the current event is over. Note that each events need to be recreated, so only the events ready to play are shown.


### Other features
*This is considered [advanced usage](https://github.com/arina999999997/elichika/blob/master/docs/advanced_usage.md)*.

Use at your own risk, the features might not even work properly on your install. 
