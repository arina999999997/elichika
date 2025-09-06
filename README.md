# elichika
A fork of https://github.com/YumeMichi/elichika, check out the original.

## Installing the game
### Android
The easiest way is to play on android is with the embedded client, which can be found [here](https://github.com/arina999999997/elichika/releases/tag/embedded).

#### Important notice
The embedded clients are only supported for 64-bits. If your device is 32-bits, you will have to use an external server or host your own server with termux.

#### Installing and playing

1. Pick the apk depending on version (Japanese or Global) and install it.
2. Open the game until the title screen (where the art of the idols are shown):
    - For Japanese version, just wait until it.
    - For Global version, select the language you want and wait until it.
3. Close the game.
4. Open the game again:
    - Some data will be extracted at this point, it will take a while.
    - The extraction is done once the title screen are shown again.
5. Once in the title screen, wait a while for the server to start up:
    - This should take no longer than 1 minute, but it depend on your device speed.
    - You can repeatedly try to start the game by tapping on the screen until it work:
      - You will repeatedly get connection error until it works.
      - However if it takes too long (>2 minutes), then something went wrong and you should seek help or try other methods.

Everytime you want to play, start the game and repeat step 5 (wait / repeatedly try).

#### Updating

As the server is still in development, it will get updates from time to time. If you want the new version with more content, you have to update. To do this:

0. Optionally, backup your data to be sure. Learn how to do it [here](https://github.com/arina999999997/elichika/blob/master/docs/webui.md).
1. Download the latest apk.
2. Install it on top of the old app:
    - Do NOT uninstall the old app, just install the new one.
3. Open the game and play like normal:
    - The first run after updates will extract some data, so you will have to wait a bit.
    - After that it's like running the game normally.

### iOS

It's a bit more involved to play on iOS, you should checkout the [LL hax wiki](https://carette.codeberg.page/ll-hax-docs/sifas/) first about how to install the app. After that, come back here if you want to host your own server.

## Documentations

### WebUI

The WebUI help you interact directly with the server outside of the game itself. Its features include (but are not limited to):

- Modify your account data:
  - Add cards, items, ...
- Create backup and restore backup:
  - This allow you to uninstall the game/server and come back later without losing progress.
  - It will also allow you to import your account into other servers (as long as it's compatible).
- Modify server behaviour:
  - Note that you need admin's right for the server to do this.
  - If you don't like certain aspect (for example, resources amount), you can change that here.
  - You can also change the active event / schedule the next event here.

To learn how to use the WebUI, checkout the [WebUI documentation](https://github.com/arina999999997/elichika/blob/master/docs/webui.md).

### Hosting your own server

If you want to host your own server (not using embedded / server hosted by other people), then you can check out the documentation [here](https://github.com/arina999999997/elichika/blob/master/docs/hosting.md). Be prepared to do work and learn how to do that though, as it might be complicated depending on how familiar with this you are, and not every steps are covered.

### Other documentations

If you are interested in learning more about the server and how it work, checkout the [documentations](https://github.com/arina999999997/elichika/blob/master/docs).

## Credit
Special thanks to the LL Hax community in general for:

- Archiving and hosting database / assets
- General and specific knowledges about the game

Even more special thanks for the specific individuals or groups (in no particular order):

- YumeMichi for original elichika.
- triangle for informations and scripts to encode/decode database, for patching the iOS clients, for daily theater logs, for databases across all versions, and for various missing asset files.
- gam for various missing asset files.
- [SIFAStheatre](https://twitter.com/SIFAStheatre) and [Idol Story](https://twitter.com/idoldotst) for Daily theater English tranlation and for the original Japanese transcript.
- ethan for hosting various resource, for hosting a public testing server, and for help with docker.
- [yunimoo](https://github.com/yunimoo) for help with docker, and for resolving TODOs.
- rayfirefirst, cppo for various cryptographic keys.
- tungnotpunk for iOS client and help with network structure.
- Suyooo for the very helpful [SIFAS wiki](https://suyo.be/sifas/wiki/), for providing more accurate stage data, and for the bad word lists.
- sarah for hosting public Internet CDN.
- AuahDark for helping with the embedded client development.
- Caret for the LL Hax discord.
- And other people who more than deserve to be here but I can't quite recall right now.

## Disclaimer
This repository is designed for official contents of SIFAS only.

The authors do not endorse any modification, except for the modifications already included in this repository. 

ALL other modifications to the assets, servers, or clients are outside of the authors' control. If such modifications cause issues, this repository is not the root cause.
