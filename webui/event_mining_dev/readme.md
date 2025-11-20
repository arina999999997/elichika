# Event mining dev
*This tool is considered [advanced usage](https://github.com/arina999999997/elichika/blob/master/docs/advanced_usage.md). You will have to setup a lot of resources file for it to work properly.*

Event mining maker to remake old events.
## Idea
- Similar to event marathon maker
## Usage

Set the config_dev.go file correctly, build the server with tag `dev` to enable the event mining dev webui.

Then start the server and go to http://127.0.0.1:8080/webui/event_mining_dev/ to start.

This might change the serverstate in unexpected way, so backup all serverside databases just to be sure.

After that just follow the routine and create the event. If working on multiple events, it might be a good idea to restart the server between making each events.

## Convention
Assume the following:
- The SR card rewarded in ranking is the one that appeared later in topic reward
- This is true for the event capture we have
## Routine
The steps to remake an event mining that the maker will follow:

- 1st: Event id - Event Name
  - This is the main decider of what the event is gonna be.
  - Since event name is stored in database already, we just pick the one linked with the event id and name.
- 2nd: Event main icon
  - The icon of the main event.
  - Select from a list of icons or manually enter the asset path.
- 3rd: Event background
  - The background art of the event.
  - Select from a list of images that usually contain the background, or manually enter the asset path:
    - The images are found by first looking in `m_story_event_history_detail` and find all `scenario_script_asset_path` associated with event id.
    - Then look inside `adv_graphic` to find the relevant asset paths.
    - Finally filter to relevant asset paths (with correct size)
- 4th: Event top still cell
  - Select from a list of cell in the same package with the event main icon, and have certain sizes.
  - Select until done.
- 5th: Event top still sub cell
  - Select from a list of commonly used subcells.
  - Usually we can't get all the cells, so there's a auto complete button that will choose from the remaining subcells:
    - The algorithm for this is to balance the amount of idols from each school.
    - Which idol is selected is random but is seeded using the event id.
- 6th: Ranking songs
  - Select 3 ranking songs and select the ranking amount (3 or 5)
- 7th: Rule descriptions pages
  - Select from a known list of commonly used rule description pages and images from the same package as main icon.
  - Select done to finish with selection and move on to next steps
- 8th: Bonus popup order
  - The order the card appear in bonus popup
  - Go through the cards that give bonus and select the position
- 9th: Point ranking topic reward
  - The order the event card show up when previewing reward accquired with ranking.
- 10th: Point ranking reward
  - Interactive builder
  - First is the template phase:
    - Select the SR order in ranking (who appear more)
    - Then select from some template:
      - Will only support 1 template to begin with, but more can be added later if necessary.
  - After the template phase, enter the adjustment phase:
    - The reward is locked, but the amount can change.
- 11th: Voltage ranking topic reward:
  - Pick from the image of the correct size on the same package as the main title.
- 12th: Voltage ranking reward:
  - Works similar to point ranking reward.
- 13th: Banner image path
  - This is used for the trade in shop
- 14th: Trade
  - This is the trade product

After going through the steps, an archive files containing the relevant data is generated.

At any point, we can go to any previous step to ammend the data:
- Each step will assume that all the steps before it are completed, and it won't use information from the later steps
- But the data for the later steps is still stored in the same object.
- Ideally, this is not necessary and should only be used to undo misinput.

## TODO
Some future stuff that are not handled totally correct yet that might be appended at the end later:
- BGM
  - The BGM will use only one song for event mining.
- Gacha associated with the event:
  - The current gacha system is written long ago without caring about non permanent gacha.
  - Then there's gacha rate up and stuff and the pool of cards
  - Then there's the short movies that show up when we load into the event.







