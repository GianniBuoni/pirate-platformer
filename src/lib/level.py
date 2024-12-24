from random import uniform
from settings import *
from lib.groups import *

from player import Player
from lib.sprites import AnimatedSprite, MovingSprite, Sprite

class Level:
    def __init__(self, tmx_map, level_frames):
        self.display_surface = pygame.display.get_surface()

        # groups
        self.all_sprites = AllSprites()
        self.collision_sprites = CollisionSprites()
        self.platform_sprites = pygame.sprite.Group()

        self.setup(tmx_map, level_frames)

    def setup(self, tmx_map, level_frames):
        # tiles
        for layer in ["BG", "Terrain", "FG", "Platforms"]:
            for x, y, surf in tmx_map.get_layer_by_name(layer).tiles():
                match layer:
                    case "BG": z = Z_LAYERS["bg tiles"]
                    case "FG": z = Z_LAYERS["fg"]
                    case _: z = Z_LAYERS["main"]

                groups: list[pygame.sprite.Group] = [self.all_sprites]
                if layer == "Terrain": groups.append(self.collision_sprites)
                if layer == "Platforms": groups.append(self.platform_sprites)

                Sprite((x * TILE_SIZE,y * TILE_SIZE), groups, surf = surf, z = z)

        # objects
        for obj in tmx_map.get_layer_by_name("Objects"):
            if obj.name == "player":
                self.player = Player((obj.x, obj.y), self.collision_sprites, self.platform_sprites, self.all_sprites)
            elif obj.name in ("crate", "barrel"): # codespell:ignore
                Sprite((obj.x, obj.y), self.all_sprites, self.collision_sprites, surf = obj.image)
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

                AnimatedSprite((obj.x, obj.y), frames, groups, z=z, animation_speed=animation_speed)

        for obj in tmx_map.get_layer_by_name("Moving Objects"):
            if obj.name == "helicopter":
                frames = level_frames[obj.name]

                if obj.width > obj.height: # horizontal movement
                    move_direction = "x"
                    start_pos = (obj.x, obj.y + obj.height / 2)
                    end_pos = (obj.x + obj.width, obj.y + obj.height / 2)
                else: # vertical movement
                    move_direction = "y"
                    start_pos = (obj.x + obj.width / 2, obj.y)
                    end_pos = (obj.x + obj.width / 2, obj.y + obj.height)
                speed = obj.properties["speed"]

                MovingSprite(start_pos, end_pos, move_direction, speed, frames, self.all_sprites, self.platform_sprites)

    def run(self, dt):
        self.display_surface.fill("black")
        self.all_sprites.update(dt)
        self.all_sprites.draw(self.player.hitbox.center)
