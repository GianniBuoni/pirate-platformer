from level import Level
from overworld import Overworld

from typing import TYPE_CHECKING
if TYPE_CHECKING: from main import Game

def switch_stage(self: "Game", target, unlock = 0):
    if target == "level":
        # Level()
        self.current_stage = Level(
            self.tmx_maps[unlock],
            self.level_frames,
            self.data,
            self.switch_stage
        )
    else:
        # Overworld
        if unlock > 0:
            self.data.unlocked_levels = (
                unlock if self.data.unlocked_levels < unlock
                else self.data.unlocked_levels
            )
        else:
            self.data.health -= 1
        self.current_stage = Overworld(
            self.tmx_overworld,
            self.overworld_frames,
            self.data,
            self.switch_stage
        )
