__all__ = ["Player"]

from settings import *
from ._enums import *

from lib.groups import CollisionSprites
from lib.timer import Timer

class Player(pygame.sprite.Sprite):
    def __init__(
            self, pos,
            frames: dict[str, list[pygame.Surface]],
            collision_sprites: CollisionSprites,
            platform_sprites,
            *groups
    ):
        super().__init__(*groups)
        self.z = Z_LAYERS["main"]

        # image
        self.frames, self.frames_idx = frames, 0
        self.facing_right = True
        self.image = self.frames[PlayerState.FALL.value][self.frames_idx]
        self.animation_speed = ANIMATION_SPEED

        # rects
        self.rect: pygame.FRect = self.image.get_frect(topleft = pos)
        self.hitbox = self.rect.inflate(-76, -36)
        self.old_rect = self.hitbox.copy()

        # collision
        self.collision_sprites = collision_sprites
        self.platform_sprites = platform_sprites

        # actions
        self.attacking = False

        # movement
        self.direction = vector()
        self.speed, self.gravity, self.jump_distance = 200, 1000, 800
        self.jump = False

        # timers
        self.timers = {
            "attack t/o": Timer(500),
            "jump t/o": Timer(400),
            "platform t/o": Timer(100),
            "wjump t/o": Timer(250),
            "damage t/o": Timer(400)
        }

    from ._animate import animate, get_state, flicker
    from ._collision import collision, platform_collision, check_collision_side, platform
    from ._core_methods import attack, input, update_timers, update
    from ._move import move
