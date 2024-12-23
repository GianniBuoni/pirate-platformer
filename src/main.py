from settings import *
from pytmx.util_pygame import load_pygame
from os.path import join

from lib.level import Level
from lib.helpers import *

class Game:
    def __init__(self):
        pygame.init()
        self.display_surface = pygame.display.set_mode((WINDOW_WIDTH, WINDOW_HEIGHT))
        pygame.display.set_caption("Super Pirate World")
        self.clock = pygame.time.Clock()
        self.import_assets()

        self.tmx_maps = {0: load_pygame(join("data", "levels", "omni.tmx"))}
        self.current_stage = Level(self.tmx_maps[0], self.level_frames)

    def import_assets(self):
        self.level_frames = {
            "big_chain": import_folder("graphics", "level", "big_chains"),
            "candle": import_folder("graphics", "level", "candle"),
            "candle_light": import_folder("graphics", "level", "candle light"),
            "flag": import_folder("graphics", "level", "flag"),
            "floor_spike": import_folder("graphics", "enemies", "floor_spikes"),
            "helicopter": import_folder("graphics", "level", "helicopter"),
            "saw": import_folder("graphics", "enemies", "saw", "animation"),
            "small_chain": import_folder("graphics", "level", "small_chains"),
            "window": import_folder("graphics", "level", "window"),
        }
        self.level_frames.update(import_sub_folders("graphics", "level", "palms"))

    def run(self):
        while True:
            dt = self.clock.tick(60) / 1000
            for event in pygame.event.get():
                if event.type == pygame.QUIT:
                    pygame.quit()
                    sys.exit()

            self.current_stage.run(dt)
            pygame.display.update()

if __name__ == "__main__":
    game = Game()
    game.run()
