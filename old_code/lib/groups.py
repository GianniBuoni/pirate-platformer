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

    def draw_overworld(self, cam_offset):
        # z sort bg
        bg_sprites = [x for x in self.sprites() if x.z < Z_LAYERS["main"]]
        for sprite in sorted(bg_sprites, key = lambda x: x.z):
            offset_pos = sprite.rect.topleft + cam_offset
            self.display_surface.blit(sprite.image, offset_pos)

        # y sort fg and player
        fg_sprites = [x for x in self.sprites() if x.z >= Z_LAYERS["main"]]
        for sprite in sorted(fg_sprites, key= lambda x: x.rect.centery):
            offset_pos = sprite.rect.topleft + cam_offset
            # offset player icon in draw method to keep rect aligned with node
            if hasattr(sprite, "player"):
                offset_pos += vector(0, -28)
            self.display_surface.blit(sprite.image, offset_pos)

class CollisionSprites(pygame.sprite.Group):
    def __init__(self, *sprites: "Sprite") -> None:
        super().__init__(*sprites)
