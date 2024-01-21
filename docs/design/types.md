# Type design
The server has to handle a lot of types, from types to talk to the clients to types to handle and store data. Therefore, we split the types into categories to make them easier to handle. The caterogies are listed in the order of how "basic" they are. Types from a category can only use types from the same category or a simpler category in its construction. 

## Client types
Client types are all the types used by the client (with exception being mentioned later on), mirrored here to read/write network data or just general handling.

These types are defined in the `client` package, and they must have the exact same construction as the type in the client. This includes the following:

- The name must match perfectly, down to the which letter are capitalised (with the exception of the first letter since we're using Go).
- The fields' types must match:

  - `int32` must be `int32` and `int64` must be `int64` and so on.
  - Replicate the type of the fields if necessary, and do not use annonymous types.
  - If the field has type `Nullable`, use the Nullable generic.
  - If the field is a pointer then it should be `Nullable` or just a value, depending on whether it can actually be `null` or not.
  
    - For example, `string` are always kept as pointer, but some strings are always filled while other can be `null` without having to be marked `Nullable`.
    - For the time being, we will use the Nullable wrapper for pointer fields that can be `null` and mark them as pointer using the comment
  - For fields that are `Dictionary`, use the `Dictionary generic`.
  - For fields that are `enum`, an enum tag to the enum name is required. This currently doesn't do anything but we might want to do enum checking and stuff later, and it just make it easier to keep track of things.
- If the type is used with `json`, it must works correctly with Marshal and Unmarshal, and it no information would be lost in doing so.

  - Use a custom marshal / unmarshaler if necessary.
  - The order of fields is not important but should still be kept.
  - The order of array elements IS important, and should be kept.
- Client types can be used by other types or used directly by handling codes, but they should not be modified to help the handling. 

  - If necessary, use an embedded type or a wrapper type.
- Finally, each type should be in its own file. The file name is derived from the type name, but we use snake_case for file names.

### Request / response types
Request and response types are also client types and follow the same rules, but they should be put into the `client/request` and `client/response` package instead. This is only done to make them easier to see. Maybe we can also do something like splitig user type into `client/user`.

Note that even if a type is only used as a subtype of a request/response type, it should be in client instead of being in `client/request` or being annonymous.

## Gamedata types
Gamedata types are types used to store how the server should work, for example, which event is there, which gacha is available, etc. These types are less constrained than the client types but should still follow:

- The naming/typing convention shoulds follow that of the same system in client's type.
- The types must be able to load from and save to database.

Gamedata types are defined in the package `gamedata` along with their loaders.

## Userdata types 
Userdata types are types used to store users' data (progress). These type should be made from Client types and Gamedata types. The general rule is as follow:

- Userdata types should have the prefix User and then the name of the relevant Client/Gamedata types.
- Userdata types should follow the Client/Gamedata type naming.
- Userdata types should contain the field UserId, of type `int`.
- Userdata types should not merge multiple data into the same table, even if that is possible:
  
  - For example, the table for `UserStatus` should not have any other info in it, even though we can store more info.
  - Or the, table for `UserMember` should only store the general, member info not things like how many card are owned.
- Userdata types should not store an information multiple times:

  - If a variable is used in multiple context, choose a single context to store it in, or store it in an entirely separate thing.
  - Aggregated data is allowed, but must be modified accordingly.

Userdata types are defined in the package `userdata` along with their loaders.

## Handling types
Handling types are types defined and used by handlers. These types can generally be just about anything. The naming convention should apply, and the handler should keep only the relevant types to itself. If a type is used a lot, it should be placed in the common utilities package instead of copied / replicated.