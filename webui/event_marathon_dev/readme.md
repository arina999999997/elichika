# Event marathon dev
Note that, this is only for remaking old event.

## Routine
The steps to remake an event marathon that the maker will follow:

- 1st: Event id
  - This is the main decider of what the event is gonna be.
  - This need to be entered manually to select what to do.
- 2nd: Event names
  - The names of the events in 4 locales.
  - Copy from various wikis for this.
- 3rd: Event main icon
  - The icon of the main event.
  - Select from a list of icons or manually enter the asset path.
- 4th: Event background
  - The background art of the event.
  - Select from a list of images that usually contain the background, or manually enter the asset path:
    - The images are found by first looking in `m_story_event_history_detail` and find all `scenario_script_asset_path` associated with event id.
    - Then look inside `adv_graphic` to find the relevant asset paths.
    - Finally filter to relevant asset paths (with correct size)
- 5th: Board texture
  - Select from a list of known texture.
  - If the known texture doesn't match, trigger a request to select from a list of texture from the same packages as previous assets.
- 6th: Board deco
  - Usually this is just null
  - If it is not null, just enter the asset path manually.
- 7th: Board memo and pictures
  - These are the notes and images on the board.
  - The order is as follow:
    - 1st memo
    - 2nd memo
    - 3rd memo
    - 1st picture
    - 2nd picture
    - ...
    - 7th picture
  - 
  - Select from images in the same package as the main icon or enter manually.
  - The priority follow this order, but if necessary, assign the priority manually later
  - Do note that we need to exit to the main screen then go back to event for this to take effect, going to menu and reloading event doesn't work
- 8th: Rule descriptions pages
  - Select from a known list of commonly used rule description pages and images from the same package as main icon.
  - Select done to finish with selection and move on to next steps
- 9th: Event booster icon
  - The icon of the event booster
  - Select from a list of icons
- 10th: Bonus popup order
  - The order the card appear in bonus popup
  - Go through the cards that give bonus and select the position
- 11th: Point rewards
  - Send a point reward file accquired through a wiki and parsed with a python script.
- 12th: Ranking rewards
  - Send a ranking reward file accquired through a wiki and parsed with a python script.
- 13th: Total topic reward
  - The order the event card show up when previewing reward accquired with points.
- 14th: Ranking topic reward
  - The order the event card show up when previewing reward accquired with ranking.

After going through the steps, an archive files containing the relevant data is generated.

At any point, we can go to any previous step to ammend the data:
- Each step will assume that all the steps before it are completed, and it won't use information from the later steps
- But the data for the later steps is still stored in the same object.
- Ideally, this is not necessary and should only be used to undo misinput.

## TODO
Some future stuff that are not handled totally correct yet that might be appended at the end later:
- BGM
  - The missing data file is not on the cdn even if we have it, so we have to host it first
- Gacha associated with the event:
  - The current gacha system is written long ago without caring about non permanent gacha.
  - Then there's gacha rate up and stuff and the pool of cards
  - Then there's the short movies that show up when we load into the event.







