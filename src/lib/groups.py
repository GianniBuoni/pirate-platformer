from settings import *

from typing import Union, TYPE_CHECKING

if TYPE_CHECKING:
    from sprites.sprites import Sprite
    from player import Player

class AllSprites(pygame.sprite.Group):
    def __init__(self, *sprites: "Union[Sprite, Player]", cam_w, cam_h, cam_t) -> None:
        super().__init__(*sprites)
        self.display_surface = pygame.display.get_surface()
        self.offset = vector()

        self.offset_limits = { # in pixels
            "left": 0,
            "right": -cam_w + WINDOW_WIDTH,
            "top": cam_t,
            "bottom": -cam_h
        }

    def constrain_camera(self):
        self.offset.x = (
            self.offset_limits["left"]
            if self.offset.x >= self.offset_limits["left"]
            else self.offset.x
        )

        self.offset.x = (
            self.offset_limits["right"]
            if self.offset.x <= self.offset_limits["right"]
            else self.offset.x
        )

        self.offset.y = (
            self.offset_limits["bottom"]
            if self.offset.y <= self.offset_limits["bottom"]
            else self.offset.y
        )

        self.offset.y = (
            self.offset_limits["top"]
            if self.offset.y >= self.offset_limits["top"]
            else self.offset.y
        )

    def draw(self, target_pos): # pyright: ignore
        self.offset.x = -(target_pos[0] - WINDOW_WIDTH / 2)
        self.offset.y = -(target_pos[1] - WINDOW_HEIGHT / 2)
        self.constrain_camera()

        for sprite in sorted(self, key = lambda x: x.z):
            offset_pos = sprite.rect.topleft + self.offset
            self.display_surface.blit(sprite.image, offset_pos)

class CollisionSprites(pygame.sprite.Group):
    def __init__(self, *sprites: "Sprite") -> None:
        super().__init__(*sprites)
