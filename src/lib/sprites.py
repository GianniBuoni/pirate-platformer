from settings import *

class Sprite(pygame.sprite.Sprite):
    def __init__(
            self, pos, surf = pygame.Surface((TILE_SIZE,TILE_SIZE)), *groups
    ) -> None:
        super().__init__(*groups)
        self.image = surf
        self.image.fill("white")
        self.rect: pygame.FRect = self.image.get_frect(topleft = pos)
        self.old_rect = self.rect.copy()

class MovingSprite(Sprite):
    def __init__(self, start_pos, end_pos, move_direction, *groups) -> None:
        super().__init__(start_pos, *groups)
