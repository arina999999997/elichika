# Hosting your own server
*Hosting your own server is considered [advanced usage](https://github.com/arina999999997/elichika/blob/master/docs/advanced_usage.md). It can be easy or hard to do depending on your specific setup, but it is not the bare minimum required to play the game, therefore it's on you to put in the work if you still want to host the server yourself.*


## Installing the server
### Android
To host the server on android, you need a shell emulator. If you have root / custom ROM, then you might be able to run it natively. Termux is recommended for this, you can get it from [f-droid](https://f-droid.org/en/packages/com.termux/) or [github](https://github.com/termux/termux-app#github). Note that the google play store version will probably works too, but there had been time when it wasn't updated for a long while and doesn't work.

There is a server installing script you can use, the following command fetch the script and run it automatically:
```
curl -L https://raw.githubusercontent.com/arina999999997/elichika/master/bin/install.sh | bash
```

### PC (Windows, Linux, MacOS)
Note that if you host the server on your PC, you will also need to make your own client, so you might want to look into that to see if you want to do it before installing the server.

You can setup the server in a desktop machine to play on android or iOS.

#### Setup manually

You can use the same script with android, but you have to install git and go first (and they have to be in your `%PATH%` for windows, and the equivalent places for Linux/MacOS). After that you can just use the same command (for windows, you also need some version of bash, git bash will works):

```
curl -L https://raw.githubusercontent.com/arina999999997/elichika/master/bin/install.sh | bash
```

Using scripts ill leave some trashes, so you can clone the respository and build manually, look at the scripts for the necessary steps.

#### Using Docker
There is a public docker image available on docker hub: https://hub.docker.com/r/arina999999997/elichika

Assuming you're familiar with docker, this can be a faster way of getting things working. Keep in mind that using docker, some of the step below will not apply, you should reference the docker docs instead.

All config options should be set in the data/config.json file, which will be created after first startup.

[docker compose](./docker/docker-compose.yml) example

## Running the server
If you host the server separately from the client, you need to run the server to play it. For Android/Linux, if you have used the script, you can use a running script:

```
~/run_elichika
```

If you have GUI for Windows/Linux, you can also just run the executable directly.

Note that whenever you want to play, the server need to be on, so if you already closed termux or the server, you will have to run it again.

## Updating the server
It is recommended to backup at least `userdata.db` before updating. You can also backup `serverstate.db`. If you have backed up your userdata separately, it's fine to not backup here. After that, you have 2 ways to update the server:

- The basic update is more stable but will take longer because it will be downloading more stuff and rebuilding more stuff.
- The normal update is faster, but if your server version is too old, something might break.

### Basic update
You can update the server using a basic update logic:

```
curl -L https://raw.githubusercontent.com/arina999999997/elichika/master/bin/basic_update.sh | bash
```

For Android/Linux where the shortcut are setup, you can also run:

```
~/basic_update_elichika
```
to do the same thing.

The basic update basically backup your data, reinstall, and then restore your data. This will work for pretty much every version, but it can be slower than the normal update.

### Normal update
If you update the server regularly, then you can use the normal update and it should work:

```
curl -L https://raw.githubusercontent.com/arina999999997/elichika/master/bin/update.sh | bash
```

If your version is new enough, then running:

```
~/update_elichika
```

would be enough.