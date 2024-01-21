# Encapsulation layers
We divide the codebase into layers of encapsulations for easier handling of things.

Each layers can contain multiple packages that can use (access) earlier layers. The orders in each layer is not defined, although it helps to keep things simple.

Here we will list the layers in the distance from the databases.

## Primitive layers
This layer contain things like defined types and configs, enum and constants.

There should be no handling logic in this part, although field taggings are allowed. For example, we can have `xorm` and `json` tags and we can have methods to modify simple things / construct things. 

## Database layers
The database layers handle initialising, reading, and writing the database.

TODO(docs): Maybe put this into its own file.
### Userdata database
Userdata database store userdata, and only userdata. It should not store things like aggregated results (ranking and such).

The table should be created using one and only one type.

- Preferablly the type should be a `client` type and should not be a merger of multiple types.
- If the interface has arrays / maps, then the arrays/maps can be split into other tables then combined when necessary
- If the arrays / maps are simple enough, it's also fine to store them just as json.

### Serverdata database
Serverdata database store server settings like what exchange are avaiable, what gacha banner is on and so on.

### Masterdata database 
This is the client database of the game. Currently `elichika` do not support modifying them, but some modification are done using another repository.

### Gamedata database
The gamedata database is the combination of the serverdata and masterdata database. It should represent the state of the games.

This database can have its own type or use `client` types.

### Cacheing database
The cacheing database store things like ranking. This database should only store things derived from the userdata database.

To make handling things simple, this database should ONLY be CALCULATED from the userdata database, with a caching invalidation policy. 

For now, there is no cache, but the following should work:
- Each data cache has a expiration time, and will be throwed away if expired.
- Cached data related to specific user expires if that user request it.


## Subsystems layer
The subsystem (name subjected to change) layer handle the subsystems of the game.

For example, handling profile, gacha or handling missions.

All the logic for handling things should be in this layer.

## Network layer
The network layer handle networking.

It should read the request and use the subsystem layer to do things.

