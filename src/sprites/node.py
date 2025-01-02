from settings import *
from .sprites import Sprite

class Node(Sprite):
    def __init__(self, pos, *groups, surf, tmx_props) -> None:
        super().__init__(pos, *groups, surf=surf, z= Z_LAYERS["path"])
        self.rect.center = (
            pos[0] + (TILE_SIZE / 2),
            pos[1] + (TILE_SIZE / 2)
        )

        self.id = tmx_props["stage"]
        # dict[input: (path_id, direction)]
        self.inputs = self.get_inputs(tmx_props)

    def get_inputs(self, tmx_props: dict[str, str]):
        inputs = {}

        for key in [x for x in tmx_props.keys() if x != "stage"]:
            if "r" in tmx_props[key]:
                path_id = int(tmx_props[key].replace("r",""))
                reverse = True
            else:
                path_id = int(tmx_props[key])
                reverse = False
            inputs[key] = (path_id, reverse)
        return inputs
