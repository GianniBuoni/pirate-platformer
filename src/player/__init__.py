from settings import *
from typing import Union

from lib.groups import CollisionSprites
from lib.sprites import MovingSprite
from lib.timer import Timer

from os.path import join

class Player(pygame.sprite.Sprite):
    def __init__(
            self, pos,
            collision_sprites: CollisionSprites,
            platform_sprites,
            *groups
    ):
        super().__init__(*groups)
        self.image = pygame.image.load(join("graphics", "player", "idle", "0.png")).convert_alpha()
        self.z = Z_LAYERS["main"]

        # rects
        self.rect: pygame.FRect = self.image.get_frect(topleft = pos)
        self.hitbox = self.rect.inflate(-76, -36)
        self.old_rect = self.hitbox.copy()

        # movement
        self.direction = vector()
        self.speed, self.gravity, self.jump_distance = 200, 1000, 800
        self.jump = False

        # collision
        self.collision_sprites = collision_sprites
        self.platform_sprites = platform_sprites
        self.platform: Union[MovingSprite, None] = None
        self.collides_with: dict[str, bool] = {
            "floor": False,
            "left": False,
            "right": False
        }

        # timers
        self.timers = {
            "jump t/o": Timer(400),
            "wjump t/o": Timer(250),
            "platform t/o": Timer(100)
        }

    from ._player_collision import collision, platform_collision, check_collision_side
    from ._player_core_methods import input, update_timers, update
    from ._player_move import move
