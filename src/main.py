from pytmx.util_pygame import load_pygame
from os.path import join

from settings import *
from lib.helpers import *
from level import Level
from lib.data import Data
from ui import UI

class Game:
    def __init__(self):
        pygame.init()
        self.display_surface = pygame.display.set_mode((WINDOW_WIDTH, WINDOW_HEIGHT))
        pygame.display.set_caption("Super Pirate World")
        self.clock = pygame.time.Clock()
        self.import_assets()

        self.ui = UI(self.font, self.ui_frames)
        self.data = Data(self.ui)
        self.tmx_maps = {0: load_pygame(join("data", "levels", "omni.tmx"))}
        self.current_stage = Level(self.tmx_maps[0], self.level_frames, self.data)

    def import_assets(self):
        self.level_frames = {
            "big_chain": import_folder("graphics", "level", "big_chains"),
            "candle": import_folder("graphics", "level", "candle"),
            "particle": import_folder("graphics", "effects", "particle"),
            "window": import_folder("graphics", "level", "window"),
            "water_top": import_folder("graphics", "level", "water", "top"),
            "water_body": import_image("graphics", "level", "water", "body"),

            # moving sprites
            "boat": import_folder("graphics", "objects", "boat"),
            "candle_light": import_folder("graphics", "level", "candle light"),
            "flag": import_folder("graphics", "level", "flag"),
            "helicopter": import_folder("graphics", "level", "helicopter"),
            "small_chain": import_folder("graphics", "level", "small_chains"),

            # enemies
            "floor_spike": import_folder("graphics", "enemies", "floor_spikes"),
            "pearl": import_image("graphics", "enemies", "bullets", "pearl"),
            "tooth": import_folder("graphics", "enemies", "tooth", "run"),
            "saw": import_folder("graphics", "enemies", "saw", "animation"),
            "saw_chain": import_image("graphics", "enemies", "saw", "saw_chain"),
            "shell": import_sub_folders("graphics", "enemies", "shell"),
            "spike": import_image("graphics", "enemies", "spike_ball", "Spiked Ball"),
            "spike_chain": import_image("graphics", "enemies", "spike_ball", "spiked_chain"),

            # player imported as a nested dictionary
            "items": import_sub_folders("graphics", "items"),
            "player": import_sub_folders("graphics", "player")
        }
        self.level_frames.update(import_sub_folders("graphics", "level", "palms"))

        self.font = pygame.font.Font(join("graphics", "ui", "runescape_uf.ttf"), 40)
        self.ui_frames = {
            "heart": import_folder("graphics", "ui", "heart"),
            "coin": import_image("graphics", "ui", "coin")
        }

    def run(self):
        while True:
            dt = self.clock.tick(60) / 1000
            for event in pygame.event.get():
                if event.type == pygame.QUIT:
                    pygame.quit()
                    sys.exit()

            self.current_stage.run(dt)
            self.ui.update(dt)
            pygame.display.update()

if __name__ == "__main__":
    game = Game()
    game.run()
