from settings import *

from typing import TYPE_CHECKING
if TYPE_CHECKING: from . import Level

def draw_skybox(self: "Level", dt):
    self.display_surface.fill("#ddc6a1")

    horizon = self.level_data["horizon_line"]
    horizon_y = horizon + self.cam_offset.y

    # water
    water_rect = pygame.FRect(0, horizon_y, WINDOW_WIDTH, horizon)
    pygame.draw.rect(self.display_surface, "#92a9ce", water_rect)
    pygame.draw.line(
        self.display_surface,
        "#f5f1de",
        (0, horizon_y),
        (WINDOW_WIDTH, horizon_y),
        4
    )

    # large clouds
    large_cloud_speed, large_cloud_dir = 50, -1
    large_cloud_tiles = int(self.level_w / self.clouds["large"].width) + 1
    self.sky_offset += large_cloud_dir * large_cloud_speed * dt
    if self.sky_offset <= -self.clouds["large"].width:
        self.sky_offset = 0

    for i in range(large_cloud_tiles):
        x = self.sky_offset + self.clouds["large"].width * i
        y = horizon_y - self.clouds["large"].height
        self.display_surface.blit(self.clouds["large"], (x, y))
