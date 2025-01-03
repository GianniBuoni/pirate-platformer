from typing import TYPE_CHECKING

from pygame import print_debug_info

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
