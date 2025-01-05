from random import uniform

from enemies.floor_spike import FloorSpike
from settings import *
from sprites import *

from typing import TYPE_CHECKING
if TYPE_CHECKING: from level import Level

def setup(self: "Level", tmx_map, level_frames):
    for obj in tmx_map.get_layer_by_name("Objects"):
        match obj.name:
            case "player":
                self.player.hitbox.topleft = (obj.x, obj.y)
            case "crate" | "barrel": # codespell:ignore
                sprites.Sprite((obj.x, obj.y), self.all_sprites, self.collision_sprites, surf = obj.image)
            case "floor_spike":
                FloorSpike(
                    (obj.x, obj.y), level_frames[obj.name],
                    True if obj.properties["inverted"] else False,
                    (self.all_sprites, self.damage_sprites),
                )
            case _:
                frames = level_frames[obj.name]
                groups: list = [self.all_sprites]
                animation_speed = ANIMATION_SPEED
                z = Z_LAYERS["main"]

                # modify frames, groups, animation_speed, and z based on object name
                if obj.name in ("palm_large", "palm_small"): groups.append(self.platform_sprites)
                if obj.name == "saw": groups.append(self.damage_sprites)
                if "palm" in obj.name: animation_speed += uniform(-1,1)
                if "bg" in obj.name: z = Z_LAYERS["bg"]

                animated.AnimatedSprite((obj.x, obj.y), frames, groups, z=z, animation_speed=animation_speed)

        if obj.name == "flag":
            self.level_flag = pygame.FRect(obj.x, obj.y, obj.width, obj.height) # pyright: ignore
