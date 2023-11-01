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

## Setting up config file
Use ``config.json`` to change some runtime parameters. The server come with a ready-to-go config by default, so if you're fine with it, you can skip this part.

- ``"cdn_server"`` 
    - The server for client to download assets.
    - The script's config use https://llsifas.catfolk.party/static/ (special thanks to sarah for hosting it).
    - We can host our own CDN with `elichika` by put the relevant files in `elichika/static`.
        - You should look into this if you want to further develop the game/server, as doing so might require redownloading things a lot.
- ``"server_address"``
    - The address to host the server at.
    - Default to ``0.0.0.0:8080`` (server listen at port 8080 on all interfaces).
- ``"tap_bond_gain"``
    - The amount of bond gained by tapping an idol.
    - Default to `20` like the original, but we can change it to a big value to skip farming bond.

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

Note that this might introduce problems because the new server might not be compatible with old database format, so you might lose progress.
 Future versions will try to keep things compatible or have a safe way to transfer. Though, if you know what you're doing, you can still transfer things over anyway.

## Playing the game
With the server running, and the client network setup correctly, simply open the game and play.

Logging in will create an account if one is not present in the server.
- User ID will be set to random if there is no user ID in the client.

### Multi accounts / Account transfer
You can use the account transfer system to switch / create account. Select ``transfer with password``. 
![Transfer system](docs/images/transfer_1.png)

Enter your user / player ID and a password:
- UserID is an non-negative integer with at most 9 digits.
- If user is in the database, password will be checked against the stored password.
- Otherwise a new account with that player ID and password.
    - You can also leave the password empty.

![Set the ID and password](docs/images/transfer_2.png)

After that, confirm the transfer and you can login with the new user ID.

![Confirm transfer](docs/images/transfer_3.png)

At any point, you can use the transfer ID system inside the game to change your password.

![Use the system](docs/images/transfer_4.png)
![Set up new password](docs/images/transfer_5.png)
![Result](docs/images/transfer_6.png)

### Client version
You can use both the Japanese and Global client for the same server (and the same database).

However, it's recommended to not play one account (user ID) in both Japanese and Global client, because some contents are exclusive to only 1 server, and will cause the client to freeze.

### Multiplayer
The current implementation doesn't explicitly support multiplayer:

- More precisely, it doesn't support 2 clients connecting at once.
- If something happens to works, it just happen to works.
- If something fails, it's the expected outcome.

So you can switch accounts and things should work, but logging in with multiple clients might result in problems.

## WebUI
The WebUI for the sever can be located at `<server_address>/webui`.
- By default, this is http://127.0.0.1:8080/webui
- The WebUI can be used to do stuff that the client can't do on its own.
    - For example, the birthday can only be set during tutorial. The WebUI can change the birthday.
- For now, the WebUI can also be used to add accessories, as drop aren't implemented properly.

## More docs
Checkout the [docs](https://github.com/arina999999997/elichika/tree/master/docs) for more details on the server and how to do more advanced stuffs. 

## Credit
Special thanks to the LL Hax community for:

- Archiving and hosting database / assets
- Original elichika release
- General and specific knowledges about the game.

