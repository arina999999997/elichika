# Elichika docs
Check out the specific documentations on how the server work and on how to do certain things.

## Server implement progress
Quick summary of what works and what doesn't.

TODO: Add specific docs for specific contents if necessary.

- [ ] Start up / New account
    - [x] Account created upon trying to login, or created using the transfer system.
    - [ ] There is no tutorial mode, but in theory it can be reconstructed.
- [ ] Login
    - [x] User can login and play.
    - [ ] There is no login bonus.
- [ ] Profile
    - [x] User can customize the profile section.
    - [ ] Birthday need to be set using the WebUI, as it is normally set during tutorial and there is no tutorial.
- [ ] Live show
    - [x] Fully working normal live, skip ticket, and 3DMV mode.
    - [x] Correctly award bond points.
    - [ ] There isn't any item drops being handled.
- [x] Story
    - [x] Fully working, you can read all kind of stories and play story songs.
    - [x] You can start from fresh and progress through the story, unlocking things that would be unlocked by story normally.
- [ ] Gacha
    - [x] Working gacha with one banner for each group.
    - [ ] Things like scouting tickets are not implemented as of now.
- [ ] Training
    - [x] Training works but always return a set of commonly used insight skills.
    - [ ] No drop from training.
- [ ] Member bond
    - [x] Working member bond system.
    - [x] Fully working bond board system.
    - [x] Bond stories are unlocked by default once you unlock the bond story feature for one member. (This might change at end of 2023 due to the database period running out).
    - [x] Bond songs unlocked at spefiic levels.
- [ ] Bond ranking
    - [x] Working but return fixed data, eventually should return actual data.
- [x] Membership
    - [x] Keep whatever membership user has in place.
    - [ ] Maybe implement a way to set the month.
- [x] Shop
    - [x] Working by returning fixed data, there is no intend to implement this further.
- [x] Exchange
    - [x] Working exchanges implementation.
    - [x] Exchange data depends on the database, by default it has items that was in the global server at the EOS.
- [x] School idol / Practice
    - [x] Fully working card grade up, level up, and practice system.
- [ ] Accessores
    - [x] Fully working accessory power up system.
    - [ ] There is no way to obtain accessories except from the WebUI, as there is no drop, and exchange for accessories doesn't work.
- [ ] Channel
    - [x] Working channel system by returning fixed data.
    - [x] User can join specific member channels.
    - [x] User can use the cheer system, but drops are not handled.
    - [ ] Megaphones are not dropped from trainings.
- [ ] Present box
    - [x] Always empty, works by returning fixed data.
    - [ ] Eventually should contains items that would be in present box instead of directly awarded.
- [ ] Goal list
    - [x] Always empty, works by returning fixed data.
    - [ ] Eventually should handle goals and daily / weekly goals.
- [x] Notices / news
    - [x] Always empty, works by returning fixed data.
    - [ ] There is no plan for now, but this section can be used to put tutorial and suchs.
- [ ] Social (friends)
    - [x] Works by returing fixed data.
    - [ ] The server allow for separate account, but implementation of social systems are not planned.
- [ ] Title List
    - [x] Stored and fetch from database.
    - [ ] There's no proper handling of adding a title yet.
- [x] Datalink
    - [x] The datalink system is used as account creation / account transfer, things should work properly.
    - [x] Password is stored in plaintext in DB because cba.
- [ ] Daily theater (JP client only)
    - [x] Works by returning fixed data.
    - [ ] Eventually should contain all the stories, user should be able to choose language some how.
    - [ ] Maybe make this feature available for WW too.
- [ ] User model
    - [x] Working user model.
    - [ ] Level up are not handled properly.

## Importing/Exporting account
You can import account from the login json or export account to json. This help with recovering your account, moving it, or update to a new server version that has a breaking change with the old database structure.

You can access this feature from the WebUI (check repository main page):

- Upload json data to import.
- Download json data for moving / backup.

Other than that, the server also generate a backup exported data everytime you login. You can find the backup in `elichika/backup`.

The import / exporting process keep 100% of the login response data, even the things that shouldn't matter by now, so you should be safe to make progress and still upgrade to the newer server version later.

For recovering data from network data (pcap), you can check out this [guide](https://github.com/arina999999997/elichika/blob/master/docs/extracting_pcap.md) on how to do that.

### How it work

- This is done using the login response from the server, which contain almost (but not quite) everything relevant to your account.
- For the information not contained in login, they are sometime can be reconstructed from context, but sometime they are just lost.

    - For example, card practice data are reconstructed from the stat of the cards given in login.

        - Note that we also only reconstruct a possible set of practice tiles, not the specific set as there could have been many.
    - Member stats on how many card they have and how many training tree filled are also reconstructed.
    - But things like how many time you used a card or how many time a card's skill was activated are not present.

        - This is avalable for at most 6 card if you have captured your profile data or have a screenshot of it, but it is just not accessible to players.

- For now we don't care that much about those data as it's not core to the gameplay experience. 

## Modifying client database

This server by default provide the databases as they were at EOS, but if necessary, you can modify the databases that the game and the server use.

This can be done to achieve the following:

- Daily songs contain all songs instead of the 3 songs per day that we have.
- Use more than 20 skip tickets at once.
- Add contents that were only in JP to WW or adding new content entirely.
- Model swap to make 12 of an idol doing a song.

You only have to modify the unencrypted database, the server will handle the rest, although it's up to you to understand the database structure and add / modify things correctly.

For more information, check [how to modify database](https://github.com/arina999999997/elichika/blob/master/docs/modify_database.md)


## How the server work

TODO: You will just have to read the code and/or ask around for now.