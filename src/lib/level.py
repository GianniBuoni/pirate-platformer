from settings import *
from lib.groups import *

from player import Player
from lib.sprites import MovingSprite, Sprite

class Level:
    def __init__(self, tmx_map):
        self.display_surface = pygame.display.get_surface()

        # groups
        self.all_sprites = AllSprites()
        self.collision_sprites = CollisionSprites()
        self.platform_sprites = pygame.sprite.Group()

        self.setup(tmx_map)

    def setup(self, tmx_map):
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

        for obj in tmx_map.get_layer_by_name("Moving Objects"):
            if obj.name == "helicopter":
                if obj.width > obj.height: # horizontal movement
                    move_direction = "x"
                    start_pos = (obj.x, obj.y + obj.height / 2)
                    end_pos = (obj.x + obj.width, obj.y + obj.height / 2)
                else: # vertical movement
                    move_direction = "y"
                    start_pos = (obj.x + obj.width / 2, obj.y)
                    end_pos = (obj.x + obj.width / 2, obj.y + obj.height)
                speed = obj.properties["speed"]
                MovingSprite(start_pos, end_pos, move_direction, speed, self.all_sprites, self.platform_sprites)

    def run(self, dt):
        self.display_surface.fill("black")
        self.all_sprites.update(dt)
        self.all_sprites.draw(self.player.hitbox.center)
