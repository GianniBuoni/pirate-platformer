from settings import *
from typing import Union

from lib.groups import CollisionSprites
from lib.sprites import MovingSprite
from lib.timer import Timer
from player._collision_enum import CollidesWith

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
        self.state, self.facing_right = "idle", True
        self.image = self.frames[self.state][self.frames_idx]
        self.animation_speed = ANIMATION_SPEED

        # rects
        self.rect: pygame.FRect = self.image.get_frect(topleft = pos)
        self.hitbox = self.rect.inflate(-76, -36)
        self.old_rect = self.hitbox.copy()

        # movement
        self.direction = vector()
        self.speed, self.gravity, self.jump_distance = 200, 1000, 800
        self.jump = False

        # actions
        self.attacking = False

        # collision
        self.collision_sprites = collision_sprites
        self.platform_sprites = platform_sprites
        self.platform: Union[MovingSprite, None] = None
        self.collides_with = CollidesWith.AIR

        # timers
        self.timers = {
            "attack t/o": Timer(500),
            "jump t/o": Timer(400),
            "platform t/o": Timer(100),
            "wjump t/o": Timer(250)
        }

    from ._player_animate import animate, get_state
    from ._player_collision import collision, platform_collision, check_collision_side
    from ._player_core_methods import attack, input, update_timers, update
    from ._player_move import move
