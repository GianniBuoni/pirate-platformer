from settings import *

class Sprite(pygame.sprite.Sprite):
    def __init__(
            self, pos, *groups,
            surf = pygame.Surface((TILE_SIZE,TILE_SIZE)),
            z = Z_LAYERS["main"],
    ) -> None:
        super().__init__(*groups)
        self.image: pygame.Surface = surf
        self.rect: pygame.FRect = self.image.get_frect(topleft = pos)
        self.old_rect = self.rect.copy()
        self.z = z
