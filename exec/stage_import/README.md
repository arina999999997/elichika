# stage_import

Import stage from json files.

## Usage
```
stage_import <path/to/file>.json
```
or 
```
stage_import <path/to/directory>
```

## Input format
Json file must be either:

- A valid json object containing the response of `"live/Start"` .
- A valid json object containing the `"live_stage"` field of the above.

## Output format
A json object containing the relevant fields:
- ``"live_difficulty_id"``: The live diffculty Id.

    - This will use the information inside the input json and not the name of the input json.
- ``"live_notes"``: The relevant fields of relevant notes. More precisely:

    - The fields `"call_time"`, `"note_type"`, `"note_position"`, `"note_action"`,  `"wave_id"`.

        - `"id"`, `"gimmick_id"`, will be renumbered when loaded.
        - `"note_random_drop_color"` is irrelevant and will be `NoteDropColor::Non (99)` for now. Even if it is used, the server will fill it in. 
        - `"auto_judge_type"` will be filled in from server config.

- ``"live_wave_settings"``: The waves. 

- ``"original"``: In the case that this stage is exactly the same as another stage, store the that here.

    - The other `"live_notes"` field will also be set to null.
    - 2 stages are considered the same if and only if the stored ``live_notes`` are exactly the same. (i.e. with irrelevant stuff filtered).
    - The original will be the one with the lowest ``"live_difficulty_id"`` from all the equal maps. 


