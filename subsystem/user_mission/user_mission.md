# How missions are handled
The missions are stored in `m_mission`.

## How missions are given
Missions are given using the `trigger_type` field, the value is one of this is `enum.MissionTrigger`:

- `enum.MissionTriggerGameStart`: Given at the start, so upon account creation / tutorial finish.

    - We will give it at tutorial finish in elichika.
    - Doesn't take any params.
- `enum.MissionTriggerClearMission`: Given upon completing another mission.

    - Take 1 param which is the mission that trigger this one.
    - We actually store this in reverse and give the new mission(s) whenever the old mission is cleared.
- `enum.MissionTriggerStoryClear`, `enum.MissionTriggerUserRank`, `enum.MissionTriggerLoveLevel`, `enum.MissionTriggerClearMusic`: Presumablly implemented but never used (checked by refencing db diff).

Asside from trigger type, there is also `pickup_type`, which affect if the mission can be given(?). This is `enum.MissionPickupType`. There are only 3 used values, aside from NULL:

- `enum.MissionPickupTypeBeginner`: For beginner goals
- `enum.MissionPickupTypeEvent`: For event goals
- `enum.MissionPickupTypeSubscription`: For subscription goals

## How missions are tracked
For each mission, we have to track the count of the goal. For daily/weekly mission, we don't reset progress but just have a starting count and ending count instead. The clear "all daily/weekly missions" seems to be reseting, howerver. Some details:

- The count will continue to go up even after the mission condition is done, but the reward is not received.
- This is probably done to not lose progress on mission chain (i.e. clear a live show with a center 100/1000/5000 times).
- Daily/weekly mission start count is tracked separately in another table, then when the mission start, we get the start_count, and we just fill in the count for all uncompleted missions.


Clear condtion (requirement type) is defined using `mission_clear_condition_type`, which is `enum.MissionClearConditionType`.

There are 103 of them, so there will be no explaination here, but their name is pretty self-planatory.

From the database, not all of them are used, so we only implement the one that are actually used.

The tracking are performed by tracking functions, one tracking function per `MissionClearConditionType`:

- The tracking functions should take the current session and the relevant params as an interface.
- The tracking functions must be called manually when needed, but there's a meta function that can call multiple functions with the same info, if it's necessary.
- The tracking functions are defined in their relevant subsystems, instead of all at once place, so they can access relevant data or functionality.
- The tracking functions are called based on change, to make it easy.

    - For example, if skip lives happen, then the tracking functions are called multiple time per function, and we can have progress added gradually instead of all at the end.
- The tracking functions should get a relevant list of mission to with the required params to track from this subsystem, or it can request mission manually.
- The mission are not sent over in delta patching, they are only sent when fetchMission is called.

