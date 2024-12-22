from settings import *
from typing import Union

from lib.groups import CollisionSprites
from lib.sprites import MovingSprite
from lib.timer import Timer

class Player(pygame.sprite.Sprite):
    def __init__(
            self, pos,
            collision_sprites: CollisionSprites,
            *groups
    ):
        super().__init__(*groups)
        self.image = pygame.Surface((48,56))
        self.image.fill("red")
        self.rect: pygame.FRect = self.image.get_frect(topleft = pos)
        self.old_rect = self.rect.copy()

        # movement
        self.direction = vector()
        self.speed, self.gravity, self.jump_distance = 200, 1000, 900
        self.jump = False

        # collision
        self.collision_sprites = collision_sprites
        self.collides_with: dict[str, bool] = {
            "floor": False,
            "left": False,
            "right": False
        }
        self.platform: Union[MovingSprite, None] = None

        # timers
        self.timers = {
            "jump t/o": Timer(400),
            "wjump t/o": Timer(250)
        }

    from ._player_collision import collision, check_collision_side
    from ._player_move import move

    def input(self):
        keys = pygame.key.get_pressed()

        if not self.timers["jump t/o"]:
            self.direction.x = int(keys[pygame.K_RIGHT]) - int(keys[pygame.K_LEFT])

        if keys[pygame.K_SPACE]:
            self.jump = True

    def update_timers(self):
        for timer in self.timers.values():
            timer.update()

    def update(self, dt):
        self.update_timers()
        self.old_rect = self.rect.copy()
        self.input()
        self.move(dt)
        self.check_collision_side()
