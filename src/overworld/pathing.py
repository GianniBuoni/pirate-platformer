from operator import ne
from settings import *
from typing import TYPE_CHECKING

from sprites.sprites import Sprite

if TYPE_CHECKING:
    from . import Overworld
    import pytmx

def get_paths(self: "Overworld", tmx_map: "pytmx.TiledMap"):
    path_points = {}
    for obj in tmx_map.get_layer_by_name("Paths"):
        obj: "pytmx.TiledObject" = obj

        if obj.properties["path_id"] <= self.data.unlocked_levels:
            for point in obj.points:
                if not obj.properties["path_id"] in path_points.keys():
                    path_points[obj.properties["path_id"]] = [(point.x, point.y)]
                else:
                    path_points[obj.properties["path_id"]].append((point.x, point.y))

    sorted_paths = {}
    for key in (sorted(path_points.keys())):
        sorted_paths[key] = path_points[key]

    self.paths = path_points

def availabe_inputs(self: "Overworld") -> tuple[list[str], dict[str, tuple[int, bool]]]: # pyright: ignore
    for node in self.node_sprites:
        if node.id == self.data.current_level:
            # trim dict to match unlocked path list
            unlocked = list(self.paths.keys())
            all_inputs = {
                x[0]: x[1] for x in node.inputs.items() if x[1][0] in unlocked
            }
            return list(all_inputs.keys()), all_inputs

def availabe_paths(self: "Overworld") -> dict[str, list[tuple[float, float]]]:
    node_info = self.availabe_inputs()
    all_paths = {}
    for key in node_info[0]:
        node_path_data = node_info[1][key]
        path = self.paths[node_path_data[0]]
        if node_path_data[1]:
            path = path[::-1]
        all_paths[key] = path
    return all_paths

def create_path(self: "Overworld", frames):
    # create tiles from path
    path_tiles = {}

    for path_id, path in self.paths.items():
        path_tiles[path_id] = []

        for i, points in enumerate(path):
            if i < len(path) - 1:
                start, end = vector(points), vector(path[i + 1])
                path_direction = (end - start) / TILE_SIZE
                start_tile = vector( # converting pixels to tile cols and rows
                    int(start[0] / TILE_SIZE), int(start[1] / TILE_SIZE)
                )

                if path_direction.y:
                    dir_y = 1 if path_direction.y > 0 else -1
                    for y in range(dir_y, int(path_direction.y) + dir_y, dir_y):
                        path_tiles[path_id].append((
                            start_tile + vector(0, y), vector(0, dir_y)
                        ))

                if path_direction.x:
                    dir_x = 1 if path_direction.x > 0 else -1
                    for x in range(dir_x, int(path_direction.x) + dir_x, dir_x):
                        path_tiles[path_id].append((
                            start_tile + vector(x, 0), vector(dir_x, 0)
                        ))

    for path in path_tiles.values():
        for i, tile_data in enumerate(path):
            tile = tile_data[0]
            direction = tile_data[1]
            other = i + 1 if i < len(path) - 1 else i - 1

            if direction == path[other][1]:
                frame_key = "vertical" if tile.x == path[other][0].x else "horizontal"
            else:
                next_vector = path[other][1]
                frame_key = {
                    (direction.x == 1 and next_vector.y == -1) or (direction.y == 1 and next_vector.x == -1): "tl",
                    (direction.x == -1 and next_vector.y == -1) or (direction.y == 1 and next_vector.x == 1): "tr",
                    (direction.x == 1 and next_vector.y == 1) or (direction.y == -1 and next_vector.x == -1): "bl",
                    (direction.x == -1 and next_vector.y == 1) or (direction.y == -1 and next_vector.x == 1): "br"
                }[True]

            Sprite((tile.x * TILE_SIZE, tile.y * TILE_SIZE), self.all_sprites, surf=frames["path"][frame_key], z=Z_LAYERS["path"])
