from lib.sprites import MovingSprite

def setup(self, tmx_map, level_frames):
    for obj in tmx_map.get_layer_by_name("Moving Objects"):
        if obj.name != "spike":
            frames = level_frames[obj.name]
            groups = (
                [self.all_sprites, self.platform_sprites] if obj.properties["platform"]
                else [self.all_sprites, self.damage_sprites]
            )

            if obj.width > obj.height: # horizontal movement
                move_direction = "x"
                start_pos = (obj.x, obj.y + obj.height / 2)
                end_pos = (obj.x + obj.width, obj.y + obj.height / 2)
            else: # vertical movement
                move_direction = "y"
                start_pos = (obj.x + obj.width / 2, obj.y)
                end_pos = (obj.x + obj.width / 2, obj.y + obj.height)
            speed = obj.properties["speed"]

            MovingSprite(start_pos, end_pos, move_direction, speed, frames, groups)
