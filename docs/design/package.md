# Packages
This file is a summary of the packages that this repository implements and what they do, and what type of package they are.

Each package might also have some docs inside of it. The local information is probably more up-to-date and accurate.
<!-- To obtain a list of package, use go list ./... -->

## Special functionality
These are the packages with functionaly outside of the common request handling stuff:

- elichika: This is the main package.
- elichika/account: implement exporting/importing an user.
- elichika/clientdb: implement the masterdata db keying process.
- elichika/encrypt: implement the necessary encryptions.
- elichika/exec/*: implement some development commands. Note that these might be out-of-date as they might depend on elichika's code that is continuing to evolve.
- elichika/klab: implement magic formula / id system that seems to be server-side only.
- elichika/utils: implement utilities code.
- elichika/webui: implement the webui system.

## Primitive layers
- elichika/config: Define config and runtime config.
- elichika/generic: Implement various generic types that the client use.
- elichika/client: Define client types mirrored here.
- elichika/client/request: Define special client types that are used for network request.
- elichika/client/response: Define special client types that are used for network response.
- elichika/enum: Define various enum values used by client, as well as some values used by the server. 
- elichika/item: Define various items/content that are commonly used by handling code. The naming of these items follow the en dictionary.

## Database layers
- elichika/dictionary: Implement the dictionary look-up for masterdata.
- elichika/serverdata: Implement and define the serverdata.db
- elichika/gamedata: Implement aggregated representation of masterdata/serverdata, so handler don't have to access the databases directly.
- elichika/userdata/database: Implement userdata database. Currently some handling code is still in elichika/userdata
- elichika/userdata: This is a mix of subsystem and database, the code should be split more cleanly. 

## Subsystem layers
- elichika/userdata: This is a mix of subsystem and database, the code should be split more cleanly. 
- elichika/subsystem/*: Implement the relevant subsystems.

    - Maybe the naming can be a bit better.

## Network layers
- elichika/router: Define the network system register point. Also define the webui endpoints for now.
- elichika/middleware: Define the common handling steps for client-server network endpoints.
- elichika/locale: Implement the locale of the request.
- elichika/handler: Declare (import) the network handlers to be used.
- elichika/handler/*: Implement and register the handler to various endpoints. The package naming is based on the url.
