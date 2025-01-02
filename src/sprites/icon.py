from enum import Enum
from settings import *

class IconStates(Enum):
    IDLE = "idle"
    LEFT = "left"
    DOWN = "down"
    UP = "up"
    RIGHT = "right"

class Icon(pygame.sprite.Sprite):
    def __init__(self, *groups, frames: dict[str, list[pygame.Surface]]) -> None:
        super().__init__(*groups)
        self.player = True

        # frames & animation
        self.state_frames, self.frame_idx = frames, 0
        self.animations_speed = ANIMATION_SPEED
        self.direction = vector()
        self.speed = 400

        # image & rect
        self.image = self.state_frames["idle"][self.frame_idx]
        self.rect: "pygame.FRect" = self.image.get_frect()
        self.z = Z_LAYERS["main"]

    def move(self, dt):
        self.rect.center += self.direction * self.speed * dt

    def animate(self, dt):
        self.frame_idx += self.animations_speed * dt
        frames = self.state_frames[self.get_state().value]
        self.image = frames[int(self.frame_idx % len(frames))]

    def get_state(self) -> IconStates:
        if self.direction == vector(-1, 0): return IconStates.LEFT
        elif self.direction == vector(1, 0): return IconStates.RIGHT
        elif self.direction == vector(0, -1): return IconStates.UP
        elif self.direction == vector(0, 1): return IconStates.DOWN
        else: return IconStates.IDLE

    def update(self, dt):
        self.move(dt)
        self.animate(dt)
