# Encapsulation layers
We divide the codebase into layers of encapsulations for easier handling of things: 

- Each layers can contain multiple packages that can use (access) earlier layers.
- The orders of package in each layer is not defined, although it helps to keep things simple.
- The layers are mostly conceptual, there is no enforced naming/grouping convention just yet. 
- You can check out the [packake listing](package.md) for a view of what each package does and which layer it fall in.

Here we will list the layers in the distance from the databases.

## Primitive layers
This layer contains definitions, like types, configs, enums, and constants.

There should be no handling logic in this part, although field taggings are allowed. For example, we can have `xorm` and `json` tags and we can have methods to modify simple things / construct things.
 

## Database layers
The database layers handle initialising, reading, and writing the database.

TODO(docs): Maybe put this into its own file.

Here is a quick summary of the types of databases:

### Userdata database
Userdata database store userdata, and only userdata. It should not store things like aggregated results (ranking and such).

The table should be created using one and only one type.

- Preferablly the type should be a `client` type plus the relevant ids like `user_id`, and should not be a merger of multiple types.
- If the interface has arrays/maps, then the arrays/maps can be split into other tables then combined when necessary.
- If the arrays/maps are simple enough, it's also fine to store them just as `json` or binary `blob`.

    - If we are always going to need the array/map field when we access the table,
    - and we are always going to need the whole of the array/map instead of just a single element,
    - then it make sense to just store the thing directly, unless it's huge.
- For now it's good enough that we can consistently save/load from the database.


### Serverdata database
Serverdata database store server settings like what exchange are avaiable, what gacha banner is on and so on. Currently these database is created directly from source code, or created from the initialisation jsons.

### Masterdata database 
This is the client database of the game. There is no reason to modify them directly inside `elichika`, although `elichika` will try to accept them modification. The modification used to setup all the feature are done using another repository.

### Gamedata database
The gamedata database is the combination of the serverdata and masterdata database. It should represent the state of the games.

This database can have its own type or use `client` types.

### Cacheing database
The cacheing database store things like ranking, other user profile, .... This database should only store things derived from the userdata database.

To make handling things simple, this database should ONLY be CALCULATED from the userdata database, with a caching invalidation policy.

For now, there is no cache, but the following polity should work in general:

- Each data cache has a expiration time, and will be recalculated if expired.
- Cached data related to specific user expires if that user request it.


## Subsystems layer
The subsystem (name subjected to change) layer handle the subsystems of the game.

For example, handling profile, gacha or handling missions.

All the logic for handling things should be in this layer.

See [subsystem docs](../../subsystem/subsystem.md) for more information. 

## Network (handler) layer
The network (handler) layer handle network request and response.

It should read the request and use the subsystem layer to do things, then generate / return the response:

- More precisely, the network layer should identify what need to be done, and use the subsystem layer to do that thing.
- The handling should depend entirely on the subsystems and not the network handler.

    - For example, clearing a live can lead to bond level up, user level up, drop, ...
    - The network handler only need to tell the subsystem to clear a live.
    - The subsystem should then handler clearing a live, returning what the network handler need and handle the conseqeuence if necessary.
    - The network handler finalize the whole process and send the response back.

Currently the code base can contain handling code in the network layer, but that should not be the case. When developing new features, we should move old handling code or write new handling code in the subsystem instead of the handling layer. 

