from settings import *
from typing import TYPE_CHECKING

if TYPE_CHECKING: from . import Player

def animate(self: "Player", dt):
    self.frames_idx += self.animation_speed * dt

    current_frames = self.frames[self.state]
    self.image = current_frames[int(self.frames_idx) % len(current_frames)]
    self.image = self.image if self.facing_right else pygame.transform.flip(self.image, True, False)

def get_state(self: "Player"):
    if self.collides_with["floor"]:
        self.state = "idle" if self.direction.x == 0 else "run"
    elif any((self.collides_with["left"], self.collides_with["right"])):
            self.state = "wall"
    else:
        self.state = "jump" if self.direction.y < 0 else "fall"
