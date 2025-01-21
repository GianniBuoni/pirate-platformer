from pytmx.util_pygame import load_pygame
from os.path import join
from typing import Union

from settings import *
from level import Level
from overworld import Overworld
from lib.data import Data
from ui import UI

class Game:
    def __init__(self):
        pygame.init()
        self.display_surface = pygame.display.set_mode((WINDOW_WIDTH, WINDOW_HEIGHT))
        pygame.display.set_caption("Super Pirate World")
        self.clock = pygame.time.Clock()

        # assets
        self.font: "Union[None, pygame.Font]" = None
        self.ui_frames, self.level_frames, self.overworld_frames, = {}, {}, {}
        self.audio = {}
        self.import_assets()

        self.ui = UI(self.font, self.ui_frames)
        self.data = Data(self.ui)
        self.tmx_maps = {
            0: load_pygame(join("data", "levels", "3.tmx")),
            1: load_pygame(join("data", "levels", "1.tmx")),
            2: load_pygame(join("data", "levels", "2.tmx")),
            3: load_pygame(join("data", "levels", "0.tmx")),
            4: load_pygame(join("data", "levels", "4.tmx")),
        }
        self.tmx_overworld = load_pygame(join("data", "overworld", "overworld.tmx"))
        self.current_stage: Union[Overworld, Level] = Overworld(
            self.tmx_overworld,
            self.overworld_frames,
            self.data,
            self.switch_stage
        )

        # events
        self.audio["starlight_city"].set_volume(.1)
        self.audio["starlight_city"].play(-1)

    from lib.import_assets import import_assets
    from lib.staging import switch_stage

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
