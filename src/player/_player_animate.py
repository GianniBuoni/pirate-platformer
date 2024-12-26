from player._collision_enum import CollidesWith
from settings import *
from typing import TYPE_CHECKING

if TYPE_CHECKING: from . import Player

def animate(self: "Player", dt):
    self.frames_idx += self.animation_speed * dt
    current_frames = self.frames[self.state]

    if self.attacking and self.frames_idx >= len(current_frames):
        current_frames = self.frames["idle"]
        self.attacking = False

    self.image = current_frames[int(self.frames_idx) % len(current_frames)]
    self.image = self.image if self.facing_right else pygame.transform.flip(self.image, True, False)

def get_state(self: "Player"):
    match self.collides_with:
        case CollidesWith.FLOOR:
            if self.attacking:
                self.state = "attack"
            else:
                self.state = "idle" if self.direction.x == 0 else "run"
        case CollidesWith.LEFT | CollidesWith.RIGHT:
            self.state = "wall"
        case CollidesWith.AIR:
            if self.attacking:
                self.state = "air_attack"
            else:
                self.state = "jump" if self.direction.y < 0 else "fall"
