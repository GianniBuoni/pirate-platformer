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
        self.jump = False
        self.jump_distance = 900

        # collision
        self.collision_sprites = collision_sprites
        self.collides_with: dict[str, bool] = {
            "floor": False,
            "left": False,
            "right": False
        }

    from ._player_collision import collision, check_collision_side

    def input(self):
        keys = pygame.key.get_pressed()
        self.direction.x = int(keys[pygame.K_RIGHT]) - int(keys[pygame.K_LEFT])

        if keys[pygame.K_SPACE] and self.collides_with["floor"]:
            self.jump = True

    def move(self, dt):
        self.rect.x += self.direction.x * self.speed * dt
        self.collision("horizontal")

        self.direction.y += self.gravity / 2 * dt
        self.rect.y += self.direction.y * dt
        self.direction.y += self.gravity / 2 * dt
        self.collision("vertical")

        # jumping logic
        if self.jump:
            self.jump = False
            self.direction.y = -self.jump_distance

    def update(self, dt):
        self.old_rect = self.rect.copy()
        self.input()
        self.move(dt)
        self.check_collision_side()
