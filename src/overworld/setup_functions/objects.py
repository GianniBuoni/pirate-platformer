from random import uniform
from settings import *
from sprites import *

from typing import TYPE_CHECKING
if TYPE_CHECKING:
    from overworld import Overworld
    import pytmx

def setup(self: "Overworld", tmx_map: "pytmx.TiledMap", frames):
    for obj in tmx_map.get_layer_by_name("Objects"):
        obj: "pytmx.TiledObject" = obj

        match obj.name:
            case "palm":
                animation_speed = ANIMATION_SPEED
                animation_speed += uniform(-1, 1)
                animated.AnimatedSprite(
                    (obj.x, obj.y),
                    frames["palms"],
                    self.all_sprites,
                    animation_speed= animation_speed,
                    z= Z_LAYERS["fg"]
                )
            case _:
                sprites.Sprite(
                    (obj.x, obj.y),
                    self.all_sprites,
                    surf = obj.image, # pyright: ignore
                )
