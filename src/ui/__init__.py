from settings import *
from .heart import Heart

class UI():
    def __init__(
        self, font,
        ui_frames: dict[str, list[pygame.Surface]]
    ) -> None:
        self.display_surface = pygame.display.get_surface()
        self.all_sprites = pygame.sprite.Group()
        self.font = font

        # health
        self.heart_frames = ui_frames["heart"]
        self.heart_w = self.heart_frames[0].get_width()
        self.heart_padding = 5
        self.create_hearts(5)

    def create_hearts(self, value):
        print("hearts")
        for i in range(value):
            x = 10 + i * (self.heart_w + self.heart_padding)
            y = 10
            Heart((x, y), self.heart_frames, self.all_sprites)

    def update(self, dt):
        self.all_sprites.update(dt)
        self.all_sprites.draw(self.display_surface)
