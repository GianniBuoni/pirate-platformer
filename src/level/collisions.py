from settings import *
from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from . import Level

def check_collisions(self: "Level"):
    # item and player
    if self.item_sprites.sprites():
        for sprite in self.item_sprites:
            collision = sprite.rect.colliderect(self.player.hitbox)
            if collision:
                sprite.kill()
                self.spawn_particle(sprite.rect.center)
                self.get_item(sprite.item_type)

    # damage sprite and player
    for sprite in self.damage_sprites:
        if hasattr(sprite, "hitbox"):
            collision = sprite.hitbox.colliderect(self.player.hitbox)
        else:
            collision = sprite.rect.colliderect(self.player.hitbox)
        if collision:
            if not self.player.timers["damage t/o"]:
                self.data.health -= 1
                self.player.timers["damage t/o"].activate()
            if hasattr(sprite, "particle"):
                sprite.kill()
                self.spawn_particle(sprite.rect.center)

    # attack collisions
    for target in self.reversible_sprites:
        collision = target.rect.colliderect(self.player.rect)
        facing_target = (
            self.player.rect.center < target.rect.center and self.player.facing_right
            or self.player.rect.center > target.rect.center and not self.player.facing_right
        )
        if collision and self.player.attacking and facing_target:
            target.reverse()

def get_item(self: "Level", item_type):
    if item_type == "potion":
        self.data.health += ITEM_VALUES["potion"]
    else:
        self.data.coins += ITEM_VALUES[item_type]
