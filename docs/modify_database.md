# Modifying database
The database the server use is at `elichika/assets/db/jp` or `elichika/assets/db/gl`, depending on the version you want to use.

The server read this database to operate, and the client download this database from the server.

However, there is a catch:

- The database in `elichika/assets/db` are in `sqlite` format.
- But the client expect an encrypted database.
- So we have to encrypt the database and send it to the client, not just the `sqlite` database

    - In theory, it's possible to modify the client to directly load the raw `sqlite`, but until that is done, we have to encrypt.

## Prequisite
- You will need `python` and will need to install some stuff:

    - You need to run `bin/encyptdbset.py`
    - To do that, install the relevant library used by that file.
    - Follow [this guide](https://discord.com/channels/922182394323292170/1102114858440327269/1137939760149704746) in LL Hax discord.
    - Maybe there will be a port of the script to Go, or the server will automatically does this everytime it's started, but not now.

## How to

To modify the database:

- First, directly modify the server's database in `elichika/assets/db/jp` or `elichika/assets/db/gl`:

    - This can be done manually using some program like [DB brower for SQLite](https://sqlitebrowser.org/)
    - It can also be done through SQL scripts, that can be executed by [DB brower for SQLite](https://sqlitebrowser.org/) or any program that support handling such scripts.
    - And you can also just replace the files with files you got from elsewhere. 

- Secondly, generate the newly encrypted database:

    ```
    ./bin/build_db.sh
    ``` 
    On Windows, you have to run inside `git bash` or some other `linux` shell emulator. 

- After that, restart the server. When you login with the client, you should be promted to download the new database.

## Notes

- This only modify the database, if you want to modify / add assets, you will also have to encrypt them. The document on doing so is not available for now.
- You might want to backup the files before trying anything.
    - The script does create backup files for you, but you will have to manually restore it.
    - The script might also fail to create a good back up.
- If your modification is inconsistent or not synced between server and client, the server or client might not work properly.
