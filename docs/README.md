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
- [ ] Story
    - [x] Fully working, you can read all kind of stories and play story songs.
    - [ ] There is no tracker for which is read and which isn't.
- [ ] Gacha
    - [x] Working gacha with one banner for each group.
    - [ ] Things like scouting tickets are not implemented as of now.
- [ ] Training
    - [ ] Training works but always return a set of commonly used insight skills.
    - [ ] No drop from training.
- [ ] Member bond
    - [x] Working member bond system.
    - [x] Fully working bond board system.
    - [ ] Bond stories are unlocked by default instead of at specific levels.
    - [ ] Bond sonds need to be unlocked by clicking the "Featured songs" button, instead of at specific levels.
- [ ] Bond ranking
    - [ ] Working but return fixed data, eventually should return actual data.
- [x] Membership
    - [x] Working by returning fixed data, there is no intend to implement this further.
- [x] Shop
    - [x] Working by returning fixed data, there is no intend to implement this further.
- [ ] Exchange
    - [x] Working exchanges implementation.
    - [x] Exchange data depends on the database, by default it has items that was in the global server at the EOS.
    - [ ] Exchange currency display aren't always correct.
- [x] School idol / Practice
    - [x] Fully working card grade up, level up, and practice system.
    - [x] Practice rewards implemented, although some might not be handled by the economy system.
    - [x] There is no material consumption for now.
- [ ] Accessores
    - [x] Fully working accessory power up system.
    - [ ] There is no way to obtain accessories except from the WebUI, as there is no drop, and exchange for accessories doesn't work.
- [ ] Channel
    - [x] Working channel system by returning fixed data.
    - [x] User can join specific member channels.
    - [ ] User can use the cheer system, but drops are not handled.
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
    - [x] Works by returning fixed data, user have all titles by default.
    - [ ] Might store in database later on.
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
    - [ ] The server keep track of items that player own pretty well, but sometime the amount is not correctly reflected.
    
## Modifying database

This server by default provide the databases as they were at EOS, but if necessary, you can modify the databases that the game and the server use.

This can be done to achieve the following:

- Daily songs contain all songs instead of the 3 songs per day that we have.
- Use more than 20 skip tickets at once.
- Add contents that were only in JP to WW or adding new content entirely.

The server provide a script to handle encoding the database for you, but (for now) it's up to you to understand the database structure and add things correctly.

For more information, check [how to modify database](https://github.com/arina999999997/elichika/blob/master/docs/modify_database.md)

## How the server work

TODO: You will just have to read the code and/or ask around for now.