from settings import *

from typing import TYPE_CHECKING
if TYPE_CHECKING: from . import Overworld

def offset_camera(self: "Overworld", target_pos: tuple):
    self.cam_offset.x = -(target_pos[0] - WINDOW_WIDTH / 2)
    self.cam_offset.y = -(target_pos[1] - WINDOW_HEIGHT / 2)
