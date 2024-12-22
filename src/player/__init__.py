from settings import *
from lib.sprites import Sprite
from lib.timer import Timer

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

        # timers
        self.timers = {
            "jump timeout": Timer(400),
            "wall jump timeout": Timer(250)
        }

    from ._player_collision import collision, check_collision_side

    def input(self):
        keys = pygame.key.get_pressed()

        if not self.timers["jump timeout"]:
            self.direction.x = int(keys[pygame.K_RIGHT]) - int(keys[pygame.K_LEFT])

        if keys[pygame.K_SPACE]:
            self.jump = True

    def move(self, dt):
        # horizontal movement
        self.rect.x += self.direction.x * self.speed * dt
        self.collision("horizontal")

        # vert
        any_wall_collide = (
            any((self.collides_with["left"], self.collides_with["right"]))
            and not self.timers["wall jump timeout"]
        )

        # wall slide decreases gravity
        if not self.collides_with["floor"] and any_wall_collide:
            self.direction.y = 0
            self.rect.y += self.gravity / 10 * dt
        else: # normal gravity
            self.direction.y += self.gravity / 2 * dt
            self.rect.y += self.direction.y * dt
            self.direction.y += self.gravity / 2 * dt

        self.collision("vertical")

        # jumping logic
        if self.jump:
            print(self.timers["wall jump timeout"].active)
            if self.collides_with["floor"]:
                self.timers["wall jump timeout"].activate()
                self.direction.y = -self.jump_distance
            elif any_wall_collide:
                self.timers["jump timeout"].activate()
                self.direction.y = -self.jump_distance
                self.direction.x = 1 if self.collides_with["left"] else -1
            self.jump = False

    def update_timers(self):
        for timer in self.timers.values():
            timer.update()

    def update(self, dt):
        self.update_timers()
        self.old_rect = self.rect.copy()
        self.input()
        self.move(dt)
        self.check_collision_side()
