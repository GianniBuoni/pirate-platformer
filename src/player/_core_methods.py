import pygame
from typing import TYPE_CHECKING

from ._enums import CollidesWith

if TYPE_CHECKING:
    from . import Player

def input(self: "Player"):
    keys = pygame.key.get_pressed()

    # regular deactivated while wall jumping
    if not self.timers["jump t/o"]:
        self.direction.x = int(keys[pygame.K_RIGHT]) - int(keys[pygame.K_LEFT])
        if self.direction.x > 0: self.facing_right = True
        if self.direction.x < 0: self.facing_right = False
        if keys[pygame.K_DOWN]:
            self.timers["platform t/o"].activate()

    # need to disable attack during wall sliding
    match self.check_collision_side():
        case CollidesWith.FLOOR | CollidesWith.AIR:
            if keys[pygame.K_f] and not self.timers["attack t/o"]:
                self.attack()

    # movement handled in jumping logic
    if keys[pygame.K_SPACE]:
        self.jump = True

def attack(self: "Player"):
    self.attacking = True
    self.frames_idx = 0
    self.timers["attack t/o"].activate()

def update_timers(self: "Player"):
    for timer in self.timers.values():
        timer.update()

def update(self: "Player", dt):
    self.update_timers()

    self.old_rect = self.hitbox.copy()
    self.input()
    self.move(dt)
    self.animate(dt)
    self.flicker()
