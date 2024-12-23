from settings import *

from typing import Union, TYPE_CHECKING

if TYPE_CHECKING:
    from lib.sprites import Sprite
    from player import Player

class AllSprites(pygame.sprite.Group):
    def __init__(self, *sprites: "Union[Sprite, Player]") -> None:
        super().__init__(*sprites)
        self.display_surface = pygame.display.get_surface()
        self.offset = vector()

    def draw(self, target_pos): # pyright: ignore
        self.offset.x = -(target_pos[0] - WINDOW_WIDTH / 2)
        self.offset.y = -(target_pos[1] - WINDOW_HEIGHT / 2)

        for sprite in self:
            offset_pos = sprite.rect.topleft + self.offset
            self.display_surface.blit(sprite.image, offset_pos)

class CollisionSprites(pygame.sprite.Group):
    def __init__(self, *sprites: "Sprite") -> None:
        super().__init__(*sprites)
