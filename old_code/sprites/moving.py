from settings import *
from .animated import AnimatedSprite

class MovingSprite(AnimatedSprite):
    def __init__(
            self, start_pos, end_pos, move_axis, speed,
            frames: list[pygame.Surface],
            *groups,
            flip = False
    ) -> None:
        super().__init__(start_pos, frames, *groups)
        self.start_pos = start_pos
        self.rect.center = self.start_pos
        self.end_pos = end_pos
        self.moving = True

        # movement props
        self.speed = speed
        self.move_axis = move_axis
        self.flip = flip
        self.direction = vector(1,0) if move_axis == "x" else vector(0,1)

    def check_flip(self):
        if self.direction.x == -1:
            self.image = pygame.transform.flip(self.image, True, False)

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

        self.animate(dt)
        if self.flip: self.check_flip()
