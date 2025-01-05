from typing import TYPE_CHECKING
if TYPE_CHECKING: from . import Level

def check_constraints(self: "Level"):
    # horizontal constraints
    if self.player.hitbox.left <= 0: self.player.hitbox.left = 0
    if self.player.hitbox.right >= self.level_w: self.player.hitbox.right = self.level_w

    # bottom constraint
    if self.player.hitbox.bottom >= self.level_h:
        self.player.direction.y = 0
        self.player.hitbox.bottom = self.level_h
        self.swith_stage("overworld", -1)

    # success
    if self.player.hitbox.colliderect(self.level_flag): # pyright: ignore
        self.swith_stage("overworld", self.level_data["level_unlock"])
