from random import uniform
from typing import TYPE_CHECKING
if TYPE_CHECKING: from . import Level

def play_sound(self: "Level", name, mod: float = 0):
    volume = uniform(0.3 + mod, 0.5 + mod)
    self.audio[name].set_volume(volume)
    self.audio[name].play()
