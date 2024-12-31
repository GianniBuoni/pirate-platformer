from settings import *
from typing import TYPE_CHECKING
if TYPE_CHECKING: from . import Level

def offset_camera(self: "Level", target_pos):
    offset_limits = { # in pixels
        "left": 0,
        "right": -self.level_w + WINDOW_WIDTH,
        "top": self.level_data["top_limit"],
        "bottom": -self.level_data["bottom_limit"]
    }

    self.cam_offset.x = -(target_pos[0] - WINDOW_WIDTH / 2)
    self.cam_offset.y = -(target_pos[1] - WINDOW_HEIGHT / 2)

    # constrain camera
    self.cam_offset.x = (
    offset_limits["left"]
    if self.cam_offset.x >= offset_limits["left"]
    else self.cam_offset.x
    )

    self.cam_offset.x = (
        offset_limits["right"]
        if self.cam_offset.x <= offset_limits["right"]
        else self.cam_offset.x
    )

    self.cam_offset.y = (
        offset_limits["bottom"]
        if self.cam_offset.y <= offset_limits["bottom"]
        else self.cam_offset.y
    )

    self.cam_offset.y = (
        offset_limits["top"]
        if self.cam_offset.y >= offset_limits["top"]
        else self.cam_offset.y
    )
