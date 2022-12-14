# alien_invasion

This repo is an implementation of the alien_invasion task.

The code was written under certain assumptions:

1. The `game_map`'s format is valid and conforms to the example specified in the task
2. Path to a city is not considered bidirectional by default. This means that if `Foo north=Boo`, then `Boo south=Foo` must be present for the path to be bidirectional
3. Upon initial distribution of aliens, if two (or more) are spawned in the same City, the City is not destroyed along with the aliens.
4. If an alien is to move from a City with no neighbours, it will remain in it and have its Travels field incremented still. This is so the game can terminate properly in an unfortunate case of a city being isolated. 