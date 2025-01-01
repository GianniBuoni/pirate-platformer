from settings import *

from typing import Union, TYPE_CHECKING

if TYPE_CHECKING:
    from sprites.sprites import Sprite
    from player import Player

class AllSprites(pygame.sprite.Group):
    def __init__(self, *sprites: "Union[Sprite, Player]") -> None:
        super().__init__(*sprites)
        self.display_surface = pygame.display.get_surface()

    def draw(self, cam_offset): # pyright: ignore
        for sprite in sorted(self, key = lambda x: x.z):
            offset_pos = sprite.rect.topleft + cam_offset
            self.display_surface.blit(sprite.image, offset_pos)

class CollisionSprites(pygame.sprite.Group):
    def __init__(self, *sprites: "Sprite") -> None:
        super().__init__(*sprites)
