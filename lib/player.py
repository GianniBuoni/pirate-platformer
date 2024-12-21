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
        self.old_rect = self.rect.copy()

        # movement
        self.direction = vector()
        self.speed = 200
        self.gravity = 1000

        # collision
        self.collision_sprites = collision_sprites

    def input(self):
        keys = pygame.key.get_pressed()
        self.direction.x = int(keys[pygame.K_RIGHT]) - int(keys[pygame.K_LEFT])
        self.direction.x = self.direction.normalize().x if self.direction else self.direction.x

    def collision(self, direction):
        for sprite in self.collision_sprites:
            if sprite.rect.colliderect(self.rect):
                if direction == "horizontal":
                    if self.rect.right >= sprite.rect.left and self.old_rect.right <= sprite.old_rect.left: # direction.x = 1
                        self.rect.right = sprite.rect.left
                    if self.rect.left <= sprite.rect.right and self.old_rect.left >= sprite.old_rect.right: # direction.x = -1
                        self.rect.left = sprite.rect.right
                else:
                    if self.rect.bottom >= sprite.rect.top and self.old_rect.bottom <= sprite.old_rect.top: # direction.y = -1
                        self.rect.bottom = sprite.rect.top
                    if self.rect.top <= sprite.rect.bottom and self.old_rect.top >= sprite.old_rect.bottom: # direction.y = 1
                        self.rect.top = sprite.rect.bottom
                    self.direction.y = 0

    def move(self, dt):
        self.rect.x += self.direction.x * self.speed * dt
        self.collision("horizontal")

        self.direction.y += self.gravity / 2 * dt
        self.rect.y += self.direction.y * dt
        self.direction.y += self.gravity / 2 * dt
        self.collision("vertical")

    def update(self, dt):
        self.old_rect = self.rect.copy()
        self.input()
        self.move(dt)
