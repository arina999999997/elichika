# Convention for handling user data.
We divide the codebase into layers of handling user data. Here we will list the layers in the distance from the userdata database.

## Userdata Database
Userdata database should be created using one and only one type.

- The type should not be merger of multiple types.
- We call type used for database `primitive` 
- If the interface has arrays / maps, then the arrays/maps can be split into other tables then combined when necessary
- If the arrays / maps are simple enough, it's also fine to store them just as json.

Userdata database should be accessed through the userdata package. The package should only handle reading/writing to the database. More precisely, it should handle reading and writing `primitive` types.

- When it's necessary to read mixed data, other layers should read the `primitive` datas and combine them.
- It's possible for other layers to store cached database as well, for example, cached ranking/profiles and such.

## Subsystems handlers
The subsystem handlers, if necessary, will handle one functionality of the game. For example, handling gacha or handling missions. This layers can help implement helper function or cache database / ....

Subsystem handlers is allowed to directly interact with the database if that will help with loading speed / ...

## Network handlers
The network handlers directly handle the network transactions. If the subsystem is simple enough, network handlers can go directly to database.
