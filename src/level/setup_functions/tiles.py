from settings import *
from sprites.sprites import Sprite

def setup_tiles(self, tmx_map, _):
    for layer in ["BG", "Terrain", "FG", "Platforms"]:
        for x, y, surf in tmx_map.get_layer_by_name(layer).tiles():
            match layer:
                case "BG": z = Z_LAYERS["bg tiles"]
                case "FG": z = Z_LAYERS["fg"]
                case _: z = Z_LAYERS["main"]

            groups: list[pygame.sprite.Group] = [self.all_sprites]
            if layer == "Terrain": groups.append(self.collision_sprites)
            if layer == "Platforms": groups.append(self.platform_sprites)

            Sprite((x * TILE_SIZE,y * TILE_SIZE), groups, surf = surf, z = z)
