from settings import *
from lib.sprites import Sprite

class CollisionSprites(pygame.sprite.Group):
    def __init__(self, *sprites: Sprite) -> None:
        super().__init__(*sprites)
