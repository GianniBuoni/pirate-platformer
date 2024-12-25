import pygame
from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from . import Player

def input(self: "Player"):
    keys = pygame.key.get_pressed()

    if not self.timers["jump t/o"]:
        self.direction.x = int(keys[pygame.K_RIGHT]) - int(keys[pygame.K_LEFT])
        if keys[pygame.K_DOWN]:
            self.timers["platform t/o"].activate()

    if keys[pygame.K_SPACE]:
        self.jump = True

def update_timers(self: "Player"):
    for timer in self.timers.values():
        timer.update()

def update(self: "Player", dt):
    self.update_timers()
    self.old_rect = self.hitbox.copy()
    self.input()
    self.move(dt)
    self.check_collision_side()
    self.animate(dt)
