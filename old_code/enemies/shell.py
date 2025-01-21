from enum import Enum

from settings import *
from lib.timer import Timer
from sprites.animated import AnimatedSprite

class ShellState(Enum):
    FIRE = "fire"
    IDLE = "idle"

class Shell(AnimatedSprite):
    def __init__(
        self, *groups, pos,
        frames: dict[str, list[pygame.Surface]],
        reverse,
        player,
        spawn_funct,
        animation_speed: float = ANIMATION_SPEED,
        z = Z_LAYERS["main"]
    ) -> None:
        if reverse: # reverse frames on init since shells are stationary
            self.state_frames = {}
            for key, surfs in frames.items():
                self.state_frames[key] = [pygame.transform.flip(x, True, False) for x in surfs]
            self.bullet_direction = -1
        else:
            self.state_frames = frames
            self.bullet_direction = 1

        super().__init__(pos, self.state_frames[ShellState.IDLE.value], *groups, animation_speed=animation_speed, z=z)

        self.player = player
        self.timeouts = {
            "fire": Timer(2000),
            "pearl spawn": Timer(1000),
        }
        self.spawn_pearl = spawn_funct

    def get_state(self) -> ShellState:
        player_pos = vector(self.player.rect.center)
        shell_pos = vector(self.rect.center)

        # distance calcs
        player_near = shell_pos.distance_to(player_pos) < 400
        player_infront = (
            shell_pos.x < player_pos.x if self.bullet_direction > 0
            else shell_pos.x > player_pos.x
        )
        player_levelwith = abs(shell_pos.y - player_pos.y) < 30

        if player_near and player_infront and player_levelwith:
            return ShellState.FIRE
        else:
            return ShellState.IDLE

    def animate(self, dt):
        self.frames = self.state_frames[self.get_state().value]

        match self.get_state():
            case ShellState.FIRE:
                if not self.timeouts["fire"]:
                    self.frames_idx = 0
                    self.timeouts["fire"].activate()
                if self.frames_idx >= len(self.frames) - 1:
                    self.frames_idx = len(self.frames) - 1

        super().animate(dt)

    def spawn_wrapper(self):
        if (
            self.get_state() == ShellState.FIRE
            and int(self.frames_idx) == 3
            and not self.timeouts["pearl spawn"]
        ):
            self.spawn_pearl(
                self.rect.midright if self.bullet_direction == 1 else self.rect.midleft,
                self.bullet_direction
            )
            self.timeouts["pearl spawn"].activate()

    def update(self, dt):
        for timer in self.timeouts.keys():
            self.timeouts[timer].update()
        super().update(dt)
        self.spawn_wrapper()
