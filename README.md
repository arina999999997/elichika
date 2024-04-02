# elichika
A fork of https://github.com/YumeMichi/elichika, check out the original.

## Installing
Note that this part concern only this server implementation. For the clients, connections, and general knowledge about the game, checkout the [LL hax wiki](https://carette.codeberg.page/ll-hax-docs/sifas/)

### Android
To run the server, first install termux, you can get it from [f-droid](https://f-droid.org/en/packages/com.termux/) or [github](https://github.com/termux/termux-app#github). Note that the google play store version will most likely NOT WORK.

Then run the install script inside termux, this will take care of everything:
```
curl -L https://raw.githubusercontent.com/arina999999997/elichika/master/bin/install.sh | bash
```
### Window / Linux
Use the same install script with termux (on Windows, run inside git bash or some other linux shell emulator):
```
curl -L https://raw.githubusercontent.com/arina999999997/elichika/master/bin/install.sh | bash
```

Or clone the respository and build manually, look at the scripts for the steps.

## Running the server
After setting up the server, we need to run it. Simply navigate to `elichika`'s directory and run it:

```
./elichika
```

For Windows:
```
elichika
```

If you have GUI for Windows/Linux, you can also just run the executable directly.

## Updating the server
You can update the server by running:

```
git pull && go build
```

As of current version, you should be able to keep your progress while doing this, so it's recommended to do it often so you get access to newly implemented features.

If you absolutely don't want to lose your data, you can always backup ``userdata.db`` and revert to an old version if necessary.

## Playing the game
With the server running, and the client network setup correctly, simply open the game and play.

Logging in will create an account if one is not present in the server.
- User id will be set to random if there is no user id in the client.

### Multi accounts / Account transfer
You can use the account transfer system to switch / create account. Select ``transfer with password``. 
![Transfer system](docs/images/transfer_1.png)

Enter your user / player id and a password:

- UserId is an non-negative integer with at most 9 digits.
- If user is in the database, password will be checked against the stored password.
- Otherwise a new account with that player id and password.
    - You can also leave the password empty.
    - If you are not running the server yourself, it's highly recommended that you setup a password, because other user can take over your account if they know your user id.
    - Passwords are securely stored with bcrypt.

![Set the id and password](docs/images/transfer_2.png)

After that, confirm the transfer and you can login with the new user id.

![Confirm transfer](docs/images/transfer_3.png)

At any point, you can use the transfer id system inside the game to change your password.

![Use the system](docs/images/transfer_4.png)
![Set up new password](docs/images/transfer_5.png)
![Result](docs/images/transfer_6.png)

### Client version
You can use both the Japanese and Global client for the same server (and the same database).

However, it's recommended to not play one account (user id) in both Japanese and Global client, because some contents are exclusive to only 1 server, and will cause the client to freeze.

### Multi devices
You can use multiple devices to play the game from one server, if you have set things up correctly.

Playing the game on another device while the current one is running will cause the current one to disconnect, preventing any error being done to your user data.

### Setting up runtime configs
There are some configuration you can change to make the server more suitable to you.

With the server running, you can use the web config editor to change runtime configs. Some changes will take effect instantly, while some will require you to restart the server.

- To access the editor, navigate to `<server_address>/config_editor` using a web browser (from the device you are playing/running the server on).

    - By default, this is http://127.0.0.1:8080/config_editor
- The config is stored in ``config.json``, so you can edit it directly if you want to.
- Howerver, it's recommended to use the web config editor as it explains some stuff and has checks to make sure nothing will go wrong.
- If your config has problems that make it so you can't run the server, you can always delete ``config.json`` to reset to the default one.

More details explanation of some config options:

- Server's address:

    - The address to host the server at.
    - Default to ``0.0.0.0:8080`` (server listen at port 8080 on all interfaces).

- CDN server's address:

    - The server for clients to download assets.
    - Default to  https://llsifas.catfolk.party/static/ (special thanks to sarah for hosting it).
    - We can host our own CDN with `elichika` by put the relevant files in `elichika/static`.
        - You should look into this if you want to further develop the game/server, as doing so might require redownloading things a lot.
    
- Default item count:

    - The amount of items to give a player to start with.
    - Default to a generous amount, but you can set it to 0 to have a more true experience.
    - Note that you have to obtain the item in game first before you are given the "default item count" amount of that item.


## WebUI
The WebUI for the sever can be located at `<server_address>/webui`.
- By default, this is http://127.0.0.1:8080/webui
- The WebUI can import and export account data.
- The WebUI can be used to do stuff that the client can't do on its own.
    - For example, the birthday can only be set during tutorial. The WebUI can change the birthday.
- The WebUI also has some account editing functions.


## More docs
Checkout the [docs](https://github.com/arina999999997/elichika/tree/master/docs) for more details on the server and how to do more advanced stuffs. 

## Credit
Special thanks to the LL Hax community in general for:

- Archiving and hosting database / assets
- General and specific knowledges about the game

Even more special thanks for the specific individuals:

- YumeMichi for original elichika.
- triangle for informations and scripts to encode/decode database, as well as patching the clients.
- rayfirefirst, cppo for various cryptographic keys.
- tungnotpunk for ios client and help with network structure.
- Suyooo for the very helpful [SIFAS wiki](https://suyo.be/sifas/wiki/) and for providing more accurate stage data.
- sarah for hosting public Internet CDN.
- Caret for the LL Hax discord.
- And other people who more than deserve to be here but I can't quite recall right now.