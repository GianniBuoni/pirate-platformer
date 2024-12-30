from settings import *
from lib.timer import Timer
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

        # coin
        self.coin_text = 0
        self.coin_timer = Timer(1000)

    def create_hearts(self, value):
        for sprite in self.all_sprites:
            sprite.kill()
        for i in range(value):
            x = 10 + i * (self.heart_w + self.heart_padding)
            y = 10
            Heart((x, y), self.heart_frames, self.all_sprites)

    def display_text(self):
        if self.coin_timer:
            text_surf = self.font.render(str(self.coin_text), False, "white")
            text_rect = text_surf.get_frect(topleft = (16, 34))
            self.display_surface.blit(text_surf, text_rect)

    def show_coin(self, value):
        self.coin_text = value
        self.coin_timer.activate()

    def update(self, dt):
        self.all_sprites.update(dt)
        self.all_sprites.draw(self.display_surface)
        self.coin_timer.update()
        self.display_text()
