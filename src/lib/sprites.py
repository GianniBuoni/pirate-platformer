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

class AnimatedSprite(Sprite):
    def __init__(
            self, pos, *groups,
            frames: list[pygame.Surface],
            animation_speed: float = ANIMATION_SPEED,
            z = Z_LAYERS["main"]
    ) -> None:
        self.frames, self.frames_idx = frames, 0
        super().__init__(pos, *groups, surf = self.frames[self.frames_idx], z=z)
        self.animation_speed = animation_speed

    def animate(self, dt):
        self.frames_idx += self.animation_speed * dt
        self.image = self.frames[int(self.frames_idx % len(self.frames))]

    def update(self, dt):
        self.animate(dt)

class MovingSprite(Sprite):
    def __init__(self, start_pos, end_pos, move_direction, speed, *groups) -> None:
        surface = pygame.Surface((200, 50))
        super().__init__(start_pos, *groups, surf = surface)
        self.image.fill("white")
        self.start_pos = start_pos
        self.rect.center = self.start_pos
        self.end_pos = end_pos
        self.moving = True

        # movement props
        self.speed = speed
        self.move_axis = move_direction
        self.direction = vector(1,0) if move_direction == "x" else vector(0,1)

    def constraints(self):
        match self.move_axis:
            case "x":
                if any((self.rect.centerx >= self.end_pos[0], self.rect.centerx <= self.start_pos[0])):
                    self.rect.centerx = self.end_pos[0] if self.direction.x == 1 else self.start_pos[0]
                    self.direction.x *= -1
            case "y":
                if any((self.rect.centery >= self.end_pos[1], self.rect.centery <= self.start_pos[1])):
                    self.rect.centery = self.end_pos[1] if self.direction.y == 1 else self.start_pos[1]
                    self.direction.y *= -1

    def update(self, dt):
        self.old_rect = self.rect.copy()
        self.rect.center += self.direction * self.speed * dt
        self.constraints()
