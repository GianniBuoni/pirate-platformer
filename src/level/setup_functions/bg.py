from settings import *
from lib.sprites import AnimatedSprite, Sprite

def setup(self, tmx_map, level_frames):
    for obj in tmx_map.get_layer_by_name("BG details"):
        z = Z_LAYERS["bg details"]
        if obj.name == "static":
            Sprite((obj.x, obj.y), self.all_sprites, surf=obj.image, z=z)
        else:
            frames = level_frames[obj.name]
            AnimatedSprite((obj.x, obj.y), frames, self.all_sprites, z=z)
            if obj.name == "candle":
                AnimatedSprite(
                    (obj.x, obj.y) + vector(-20, -20),
                    level_frames["candle_light"],
                    self.all_sprites,
                    z=z
                )
