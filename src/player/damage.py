from math import sin
from settings import *

from typing import TYPE_CHECKING
if TYPE_CHECKING: from . import Player

def damaged(self: "Player"):
    if not self.timers["damage t/o"]:
        print("player hit")
        self.timers["damage t/o"].activate()

def flicker(self: "Player"):
    if self.timers["damage t/o"] and self.image and sin(pygame.time.get_ticks() / 50) >= 0:
        white_mask = pygame.mask.from_surface(self.image)
        white_surface = white_mask.to_surface()
        white_surface.set_colorkey("black")
        self.image = white_surface
