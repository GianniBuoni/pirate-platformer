from ._enums import CollidesWith, PlayerState
from settings import *
from typing import TYPE_CHECKING

if TYPE_CHECKING: from . import Player

def animate(self: "Player", dt):
    self.frames_idx += self.animation_speed * dt
    current_frames = self.frames[self.get_state().value]

    if self.attacking and self.frames_idx >= len(current_frames):
        current_frames = self.frames["idle"]
        self.attacking = False

    self.image = current_frames[int(self.frames_idx) % len(current_frames)]
    self.image = self.image if self.facing_right else pygame.transform.flip(self.image, True, False)

def get_state(self: "Player") -> PlayerState:
    match self.check_collision_side():
        case CollidesWith.FLOOR:
            if self.attacking:
                return PlayerState.ATTACK
            else:
                return PlayerState.IDLE if self.direction.x == 0 else PlayerState.RUN
        case CollidesWith.LEFT | CollidesWith.RIGHT:
            return PlayerState.WALL
        case CollidesWith.AIR:
            if self.attacking:
                return PlayerState.AIR_ATK
            else:
                return PlayerState.JUMP if self.direction.y < 0 else PlayerState.FALL
