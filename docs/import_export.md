# Importing/Exporting account
You can use the user WebUI to export and import your account.

You can use either `json` or `db` formats.

## DB
This is the prefered format for this implementation. It's basically a database that contain only 1 account. It save every data, and exporting and reimporting should result in no change in the account. Howerver, there are some limitation:

- The friend data is not extracted, as the friends are not guaranteed to exists in whatever server you import to:
    
    - More precisely, the friends are server side only, and doesn't change no matter what account you import.
    - Even if you import another account to your current user id, you still have the same friend set.
    - But importing your exported account to another server will use that server's friend set.

- Credential data is also not extracted in a similar manner.

## JSON
You can import account from the login json or export account to json. This help with recovering your account, moving it, or update to a new server version (or another different server).

The server also generate a backup exported data everytime you login. You can find the backup in `elichika/backup` on the server machine.

This can also be used to recover data from captured network data (pcap), you can check out this [guide](https://github.com/arina999999997/elichika/blob/master/docs/extracting_pcap.md) on how to do that.

### How it work

- This is done using the login response from the server, which contain almost (but not quite) everything relevant to your account.
- For the information not contained in login, they are sometime can be reconstructed from context, but sometime they are just lost:

    - For example, card practice data are reconstructed from the stat of the cards given in login.

        - Note that we also only reconstruct a possible set of practice tiles, not the specific set as there could have been many.
    - Member stats on how many card they have and how many training tree filled are also reconstructed.
    - But things like how many time you used a card or how many time a card's skill was activated are not present.

        - This is avalable for at most 6 card if you have captured your profile data or have a screenshot of it, but it is just not accessible to players.

- For now we don't care that much about those data as it's not core to the gameplay experience. 
