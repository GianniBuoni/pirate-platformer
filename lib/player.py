from settings import *
from lib.sprites import Sprite

class Player(pygame.sprite.Sprite):
    def __init__(
            self, pos,
            collision_sprites: "pygame.sprite.Group[Sprite]",
            *groups
    ):
        super().__init__(*groups)
        self.image = pygame.Surface((48,56))
        self.image.fill("red")
        self.rect: pygame.FRect = self.image.get_frect(topleft = pos)

    # movement
        self.direction = vector()
        self.speed = 200

    # collision
        self.old_rect = self.rect.copy()
        self.collision_sprites = collision_sprites

    def input(self):
        keys = pygame.key.get_pressed()
        self.direction.x = int(keys[pygame.K_RIGHT]) - int(keys[pygame.K_LEFT])
        self.direction = self.direction.normalize() if self.direction else self.direction

    def collision(self):
        for sprite in self.collision_sprites:
            # collide on x
            if self.rect.right >= sprite.rect.left and self.old_rect.right <= sprite.old_rect.left:
                pass
            if self.rect.left <= sprite.rect.right and self.old_rect.left >= sprite.old_rect.left:
                pass

    def move(self, dt):
        self.rect.topleft += self.direction * self.speed * dt

    def update(self, dt):
        self.old_rect = self.rect.copy()
        self.input()
        self.move(dt)
