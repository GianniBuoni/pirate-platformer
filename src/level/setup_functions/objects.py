from random import uniform

from settings import *
from sprites import *
from player import Player

def setup(self, tmx_map, level_frames):
        # objects
    for obj in tmx_map.get_layer_by_name("Objects"):
        if obj.name == "player":
            frames = level_frames[obj.name]
            self.player = Player((obj.x, obj.y), frames, self.collision_sprites, self.platform_sprites, self.all_sprites)
        elif obj.name in ("crate", "barrel"): # codespell:ignore
            sprites.Sprite((obj.x, obj.y), self.all_sprites, self.collision_sprites, surf = obj.image)
        else: # all other animated sprites on the objects layer
            frames = level_frames[obj.name]
            groups = [self.all_sprites]
            animation_speed = ANIMATION_SPEED
            z = Z_LAYERS["main"]

            # modify frames, groups, animation_speed, and z based on object name
            if obj.name == "floor_spike" and obj.inverted:
                frames = [pygame.transform.flip(x, False, True) for x in frames]
            if obj.name in ("palm_large", "palm_small"): groups.append(self.platform_sprites)
            if "palm" in obj.name: animation_speed += uniform(-1,1)
            if "bg" in obj.name: z = Z_LAYERS["bg"]

            animated.AnimatedSprite((obj.x, obj.y), frames, groups, z=z, animation_speed=animation_speed)
