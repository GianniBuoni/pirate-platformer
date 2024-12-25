from typing import TYPE_CHECKING

if TYPE_CHECKING: from . import Player

def animate(self: "Player", dt):
    self.frames_idx += self.animation_speed * dt

    current_frames = self.frames[self.state]
    self.image = current_frames[int(self.frames_idx) % len(current_frames)]
