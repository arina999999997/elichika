# Modifying database
*This is considered [advanced usage](https://github.com/arina999999997/elichika/blob/master/docs/advanced_usage.md)*

## Why?
This server by default provide the databases as they were at EOS, plus the relevant modification for the features (turning DLP on and so). If you wanted to, you can modify the databases that the game and the server use.

This can be done to achieve the following, but also much more:

- Daily songs contain all songs instead of the 3 songs per day that we have.
- Use more than 20 skip tickets at once.
- Add contents that were only in JP to WW or adding new content entirely.
- Model swap to make 12 of an idol doing a song.

You only have to modify the unencrypted database, the server will handle the rest, although it's up to you to understand the database structure and add / modify things correctly.

## Something to keep in mind
The database the server use is at `elichika/assets/db/jp` or `elichika/assets/db/gl`, depending on the version you want to use.

The server read this database to operate, and the client download this database from the server.

However, there is a catch:

- The database in `elichika/assets/db` are in `sqlite` format.
- But the client expect an encrypted database.
- So we have to encrypt the database and send it to the client, not just the `sqlite` database

    - In theory, it's possible to modify the client to directly load the raw `sqlite`, but until that is done, we have to encrypt.

## How to modify database with elichika?

To modify the database, directly modify the server's database in `elichika/assets/db/jp` or `elichika/assets/db/gl`:

- This can be done manually using some program like [DB brower for SQLite](https://sqlitebrowser.org/)
- It can also be done through SQL scripts, that can be executed by [DB brower for SQLite](https://sqlitebrowser.org/) or any program that support handling such scripts.
- And you can also just replace the files with files you got from elsewhere.

After that, you only need to restart the server and it will automatically generate the necessary files. Then you only have to login or move around with the client to trigger a database update on client side.

## Notes

- This only modify the database, if you want to modify / add assets, you will also have to encrypt them. The document on doing so is not available for now.
- You might want to backup the files before trying anything.
- If your modification is inconsistent or not synced between server and client, the server or client might not work properly.
- Ideally, you should save the modification you made as SQL scripts, so you can repeat them whenever you want with a fresh database instance.
- New version of elichika WILL make change to the database, so having a modification and updating will not work properly. You have to update first, then apply your modification again.