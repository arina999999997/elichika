# Member guild (channel)
All member guild data are stored inside `u_member_guild` and `u_member_guild_support_item`.

Additionally, there is the `u_mmember_guild_coop_point` table that is derived from `u_member_guild`, this table is calculated on necessary and reseting it can lead to wrong state, but nothing critical.

## Server locale rule
Ranking and rewards are calculated based on user's (client's) locale instead of being fixed on the server:

- Whnever ranking is calculated, all users are considered no matter if they play on gl or jp.
- If the user is using the gl client, gl rule for rewards calculation apply.
- If the user is using the jp client, jp rule for rewards calculation apply.

## Constraint on periods
For simplicity, we assume the following w.r.t the period:

- The OneCycleSecs (duration of each cycle) must be a multiple of 1 days.
- 0 = TransferStartSecs < TransferEndSecs = RankingStartSecs < RankingEndSecs < OneCycleSecs

We also assume the following:

- There are 2 free cheer per day, they are given out at start of day (00:00) and mid day (12:00)
- The daily tracking for point reset at noon (middle) of day.

## Coop point
Daily coop point calculation:

- The daily coop point of current day is calculated directly using the `u_member_guild` table:

  - Just sum up the relevant point of people with the same reset time
- The daily coop point of previous day is calculated once when someone trigger the new day reset, otherwise it is taken from the table `m_member_guild_daily_coop`:

    - Storing this is memory is also an option, as long as we assume that there won't be 2 restarts of the server in 1 day.

## Rewards
### Cheer reward
Cheer reward are sent to present box upon cheering.

### Final ranking reward
Rewards are triggered upon bootstraping after the ranking:

- The reward for a member guild can be received at anytime during the next member guild, but has to be triggered to be received.
- The server will use the database `u_member_guild_receive_reward` to track 
- It might be possible to trigger reward on loading top too


### Daily coop reward
Daily coop reward are delivered upon reset, but loading the channel menu is required to actually get them, people who don't load the channel menu will miss out on it even if they actually had contribution. 

## Internal
The total point field isn't always up-to-date in code, it will be correct in database however. 