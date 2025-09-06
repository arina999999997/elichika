# Elichika docs
Check out the specific documentations on how the server work and on how to do certain things.

## Server implement progress
Quick summary of what works and what doesn't. If you want to see it in a more technical sense, checkout the [endpoints](https://github.com/arina999999997/elichika/blob/master/router/endpoints.md) listing.

TODO(docs): Add specific docs for specific contents if necessary.

- [x] Start up / New account
    - [x] Account created upon trying to login, or created using the transfer system.
    - [x] New account will trigger the openning MV and the tutorial mode. The tutorial process all work although some part can be improved.
- [x] Login
    - [x] User can login and play.
    - [x] There are login bonus types like normal login bonus, idol birthday login bonus, and new player login bonus.
- [x] Profile
    - [x] User can customize the profile section.
    - [x] Birthday can be set during tutorial or changed using the WebUI.
- [x] Live show
    - [x] Fully working normal live, skip ticket, and 3DMV mode.
    - [x] Correctly award bond points.
    - [x] You can use your own partner guests. 
    - [x] Drops are handled "correctly"
- [x] Story
    - [x] Fully working, you can read all kind of stories and play story songs.
    - [x] You can start from fresh and progress through the story, unlocking things that would be unlocked by story normally.
- [ ] Gacha
    - [x] Working gacha with one banner for each group.
    - [ ] Things like scouting tickets are not implemented as of now.
- [x] Training
    - [x] Training works but always return a set of commonly used insight skills.
    - [x] Training drops items, and drops rally megaphone while channel live is on.
- [x] Member bond
    - [x] Working member bond system.
    - [x] Fully working bond board system.
    - [x] Bond stories are unlocked by level once you unlock the bond story feature for one member (get to level 3 bond).
    - [x] Bond songs unlocked at spefiic levels.
- [x] Bond ranking
    - [x] Working bond ranking, but it might be slow if there are a lot of account.
- [ ] Membership (subscription)
    - [x] Keep membership info for imported data.
    - [x] Add default membership for new account.
    - [ ] There is no tracking or veteran reward.
- [x] Shop
    - [x] Working by returning fixed data, there is no intend to implement this further.
- [x] Exchange
    - [x] Working exchanges implementation.
    - [x] Exchange data depends on the database, by default it has items that was in the global server at the EOS.
- [x] School idol / Practice
    - [x] Fully working card grade up, level up, and practice system.
- [x] Accessories
    - [x] Fully working accessory power up system.
    - [x] Accesory drops from live and can be exchanged in shop.
    - [x] The WebUi functionality to add accessory is still there for accessory that are limited or can't be dropped, if you wish to get them
- [x] Channel
    - [x] Working channel system with ranking reward and reward.
- [x] Present box
    - [x] Working present box.
    - [x] All items that are sent to present box should be there, but there might be mistakes.
- [x] Goal list
    - [x] Working daily / weekly goals that reset correctly.
    - [x] Working goals tracking for free goals that are available at EOS
    - [ ] Some other event exclusive goals are not implemented, they might be revived later on.
- [x] Notices / news
    - [x] Always empty, works by returning fixed data.
    - [ ] There is no plan for now, but this section can be used to put tutorial and suchs.
- [x] Social (friends)
    - [x] Working social system.
    - [x] Working bad word checker.
- [x] Title List
    - [x] Stored and fetch from database.
    - [x] Title is added through user content system.
    - [x] User can obtain title through goals and suchs
- [x] Datalink
    - [x] The datalink system is used as account creation / account transfer, things should work properly.
    - [x] Password is stored using bcrypt, so no worry of leaking password. 
- [x] Daily theater
    - [x] Working daily theater server code. 
    - [x] Working Global client with the feature enabled.
    - [x] Japanese text use network log or transcript, English text use translation (thanks to [SIFAStheatre](https://twitter.com/SIFAStheatre) and [Idol Story](https://twitter.com/idoldotst))
    - [ ] Korean and Chinese (zh) translation is not avaialble.
- [x] User model
    - [x] Working user model.
    - [x] Working LP and AP recovery system (in original resource setting)
- [x] DLP
    - [x] Working DLP that also track voltage ranking
    - [x] User can reset DLP progress using the WebUi.
- [ ] Event
    - [x] Working marathon (story/point reward) event handling, with the first event available.
    - [ ] Event goals / gacha are not available.
    - [ ] Other marathon events also have missing assets that need to be remade.
    - [ ] There are also some other limitation/defect due to the current design.
    - [ ] Mining (exchange) event.
    - [ ] SBL event.
    - [ ] Voltage ranking event.
    
## Importing/Exporting account
You can use the user WebUI to export and import your account.

You can use either `json` or `db` formats.

### DB
This is the prefered format for this implementation. It's basically a database that contain only 1 account. It save every data, and exporting and reimporting should result in no change in the account. Howerver, there are some limitation:

- The friend data is not extracted, as the friends are not guaranteed to exists in whatever server you import to:
    
    - More precisely, the friends are server side only, and doesn't change no matter what account you import.
    - Even if you import another account to your current user id, you still have the same friend set.
    - But importing your exported account to another server will use that server's friend set.

- Credential data is also not extracted in a similar manner.

### JSON
You can import account from the login json or export account to json. This help with recovering your account, moving it, or update to a new server version (or another different server).

The server also generate a backup exported data everytime you login. You can find the backup in `elichika/backup`.

This can also be used to recover data from captured network data (pcap), you can check out this [guide](https://github.com/arina999999997/elichika/blob/master/docs/extracting_pcap.md) on how to do that.

#### How it work

- This is done using the login response from the server, which contain almost (but not quite) everything relevant to your account.
- For the information not contained in login, they are sometime can be reconstructed from context, but sometime they are just lost:

    - For example, card practice data are reconstructed from the stat of the cards given in login.

        - Note that we also only reconstruct a possible set of practice tiles, not the specific set as there could have been many.
    - Member stats on how many card they have and how many training tree filled are also reconstructed.
    - But things like how many time you used a card or how many time a card's skill was activated are not present.

        - This is avalable for at most 6 card if you have captured your profile data or have a screenshot of it, but it is just not accessible to players.

- For now we don't care that much about those data as it's not core to the gameplay experience. 

## Modifying client database
*This is considered [advanced usage](https://github.com/arina999999997/elichika/blob/master/docs/advanced_usage.md)*

This server by default provide the databases as they were at EOS, plus the relevant modification for the features (turning DLP on and so), but if necessary, you can modify the databases that the game and the server use.

This can be done to achieve the following, but also much more:

- Daily songs contain all songs instead of the 3 songs per day that we have.
- Use more than 20 skip tickets at once.
- Add contents that were only in JP to WW or adding new content entirely.
- Model swap to make 12 of an idol doing a song.

You only have to modify the unencrypted database, the server will handle the rest, although it's up to you to understand the database structure and add / modify things correctly.

For more information, check [how to modify database](https://github.com/arina999999997/elichika/blob/master/docs/modify_database.md)


## How the server work

You will just have to read the code if you want more details.