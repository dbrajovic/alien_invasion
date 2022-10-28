# alien_invasion

This repo is an implementation of the alien_invasion task.

The code was written under certain assumptions:

1. The `game_map`'s format is valid and conforms to the example specified in the task
2. Path to a city is not considered bidirectional by default. This means that if `Foo north=Boo`, then `Boo south=Foo` must be present for the path to be bidirectional
3. Upon initial distribution of aliens, if two (or more) are spawned in the same City, the City is not destroyed along with the aliens.