from settings import *

from typing import TYPE_CHECKING
if TYPE_CHECKING: from . import Overworld

def offset_camera(self: "Overworld", target_pos: tuple):
    self.cam_offset.x = -(target_pos[0] - WINDOW_WIDTH / 2)
    self.cam_offset.y = -(target_pos[1] - WINDOW_HEIGHT / 2)

def move_icon(self: "Overworld"):
    if len(self.current_path) == 0:
        self.icon.direction = vector()
        self.start_point = None
    else:
        destination_point = self.current_path[0]
        if self.start_point:
            if self.start_point[0] > destination_point[0]: # move left
                self.icon.direction.x = -1 if self.icon.rect.left > destination_point[0] else 0
                self.pivot_path_points(self.icon.direction.x)
            if self.start_point[0] < destination_point[0]: # move right
                self.icon.direction.x = 1 if self.icon.rect.left < destination_point[0] else 0
                self.pivot_path_points(self.icon.direction.x)
            if self.start_point[1] > destination_point[1]: # move up
                self.icon.direction.y = -1 if self.icon.rect.top > destination_point[1] else 0
                self.pivot_path_points(self.icon.direction.y)
            if self.start_point[1] < destination_point[1]: # move down
                self.icon.direction.y = 1 if self.icon.rect.top < destination_point[1] else 0
                self.pivot_path_points(self.icon.direction.y)

def pivot_path_points(self: "Overworld", speed):
    if speed == 0:
        self.icon.rect.topleft = self.current_path[0]
        self.start_point = self.icon.rect.topleft
        del self.current_path[0]
        self.check_current_node()
