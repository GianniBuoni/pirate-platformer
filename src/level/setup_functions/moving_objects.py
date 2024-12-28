from settings import *
from lib.sprites import Sprite
from lib.moving_sprites import MovingSprite

def setup(self, tmx_map, level_frames):
    for obj in tmx_map.get_layer_by_name("Moving Objects"):
        if obj.name == "spike":
            pass
        else:
            frames = level_frames[obj.name]
            groups = (
                [self.all_sprites, self.platform_sprites] if obj.properties["platform"]
                else [self.all_sprites, self.damage_sprites]
            )
            flip = obj.properties["flip"]

            if obj.width > obj.height: # horizontal movement
                move_direction = "x"
                start_pos = (obj.x, obj.y + obj.height / 2)
                end_pos = (obj.x + obj.width, obj.y + obj.height / 2)
            else: # vertical movement
                move_direction = "y"
                start_pos = (obj.x + obj.width / 2, obj.y)
                end_pos = (obj.x + obj.width / 2, obj.y + obj.height)
            speed = obj.properties["speed"]

            MovingSprite(start_pos, end_pos, move_direction, speed, frames, groups, flip = flip)

            # draw saw chains for saw objects
            if obj.name == "saw":
                surfaces = level_frames["saw_chain"]
                draw_saw_chains(self, surfaces, start_pos, end_pos, move_direction)

def draw_saw_chains(self, surfaces, start_pos, end_pos, move_direction):
    if move_direction == "x":
        y = start_pos[1] - surfaces.get_height() / 2
        left, right = int(start_pos[0]), int(end_pos[0])
        for x in range(left, right, 20):
            Sprite((x,y), self.all_sprites, surf=surfaces, z=Z_LAYERS["bg details"])

    else:
        x = start_pos[0] - surfaces.get_width() / 2
        top, bottom = int(start_pos[1]), int(end_pos[1])
        for y in range(top, bottom, 20):
            Sprite((x,y), self.all_sprites, surf=surfaces, z=Z_LAYERS["bg details"])
