from settings import *
from typing import TYPE_CHECKING

if TYPE_CHECKING: from overworld import Overworld

def input(self: "Overworld"):
    keys = pygame.key.get_pressed()
    valid_inputs = self.availabe_inputs()[0]

    def go(direction):
        self.start_point = self.icon.rect.topleft
        self.current_path = self.availabe_paths()[direction][1:]
        self.can_input = False

    if self.can_input and not self.timers["input t/o"]:
        if keys[pygame.K_LEFT] and "left" in valid_inputs: go("left")
        if keys[pygame.K_RIGHT] and "right" in valid_inputs: go("right")
        if keys[pygame.K_UP] and "up" in valid_inputs: go("up")
        if keys[pygame.K_DOWN] and "down" in valid_inputs: go("down")
        if keys[pygame.K_RETURN]:
            self.switch_stage("level", self.data.current_level)
