# Modifying database
The database the server use is at `elichika/assets/db/jp` or `elichika/assets/db/gl`, depending on the version you want to use.

The server read this database to operate, and the client download this database from the server.

However, there is a catch:

- The database in `elichika/assets/db` are in `sqlite` format.
- But the client expect an encrypted database.
- So we have to encrypt the database and send it to the client, not just the `sqlite` database

    - In theory, it's possible to modify the client to directly load the raw `sqlite`, but until that is done, we have to encrypt.

## How to

To modify the database, directly modify the server's database in `elichika/assets/db/jp` or `elichika/assets/db/gl`:

    - This can be done manually using some program like [DB brower for SQLite](https://sqlitebrowser.org/)
    - It can also be done through SQL scripts, that can be executed by [DB brower for SQLite](https://sqlitebrowser.org/) or any program that support handling such scripts.
    - And you can also just replace the files with files you got from elsewhere.
    - TODO: Add database mod to WebUI?

After that, you only need to restart the server and it will automatically generate the necessary files. Then you only have to login or move around with the client to trigger a database update on client side.

## Notes

- This only modify the database, if you want to modify / add assets, you will also have to encrypt them. The document on doing so is not available for now.
- You might want to backup the files before trying anything.
- If your modification is inconsistent or not synced between server and client, the server or client might not work properly.
- Ideally, you should save the modifcation you made as SQL scripts, so you can repeat them whenever you want with a fresh database instance.