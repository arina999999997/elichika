# Subsystem
Subsystem handles various part of a specific system in the game, and the consequence of such actions.

For example, whenever we finish a live, we need to add experience to the user. This will potentially trigger a rank up, that will trigger a rank up reward.

Or whenever we finish gacha, we need to add the card, which will in turn level up bond limit and/or give grade up items. Leveluping bond limit can also unlock new bond board and such.

## Consequence list 
This consequence list is not completed. Hopefully there's no loop in the subsystem or we might have to have one big package or split a subsystem into multiple parts.

```
UserContent <- Live (exp/drop)
UserContent <- Card (practice)
UserContent <- Member (bond reward)
Member <- Live (love)
Member <- Card (love limit)
```