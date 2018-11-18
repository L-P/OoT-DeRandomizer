# OoT-DeRandomizer
Extract a spoiler log from a ROM generated by the [OoT Randomizer][1].  
This is essentially a cheat, never use it in a race.

[1]: https://github.com/AmazingAmpharos/OoT-Randomizer

## Why
There was some cheating accusations thrown around in the OoT randomizer race
community and I wanted to prove it was feasible and easy to cheat, if such a
program publicly exists there can be no doubt about the existence of cheaters.  

## How
By reading back the item override list from the ROM you can match locations and
items. There is more to look for but the resulting spoiler log is quite
complete already.

## How to prevent cheating
I can see two solutions:

1. Don't let the runner handle the ROM, someone else should set up the hardware
   and load the ROM. The ROM should be generated seconds before the race and
   not released publicly.  
   This is only viable in a IRL tournament setting.

2. Encrypt the item and hint lists, let the player input a decryption key
   before starting the run. This is viable for online races where you can
   distribute the ROM beforehand and the key right before the run starts.  
   This is the most efficient method that prevents offline ROM scanning.

## Bugs
1. Some items randomly miss
2. Some locations randomly give wrong items such as GS appearing without
   skullsanity.
3. MQ is poorly handled, if a single dungeon is MQ all MQ locations will
   appear.
4. Medallions are missing.
