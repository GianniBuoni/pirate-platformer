from settings import *
from enemies import *

from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from level import Level

def setup(self: "Level", tmx_map, level_frames):
    for obj in tmx_map.get_layer_by_name("Enemies"):
        frames = level_frames[obj.name]
        groups: list = [self.all_sprites]

        match obj.name:
            case "tooth":
                groups.append(self.damage_sprites)
                tooth.Tooth((obj.x, obj.y), frames, self.collision_sprites, groups)
            case "shell":
                shell.Shell(
                    groups,
                    pos = (obj.x, obj.y),
                    frames = frames,
                    reverse = obj.properties["reverse"],
                    player = self.player,
                    spawn_funct = self.spawn_pearl
                )
