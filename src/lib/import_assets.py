from lib.helpers import *

from typing import TYPE_CHECKING
if TYPE_CHECKING: from main import Game

def import_assets(self: "Game"):
    self.level_frames = {
        "big_chain": import_folder("graphics", "level", "big_chains"),
        "candle": import_folder("graphics", "level", "candle"),
        "cloud_large": import_image("graphics", "level", "clouds", "large_cloud"),
        "cloud_small": import_folder("graphics", "level", "clouds", "small"),
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

        # imported as a nested dictionary
        "bg": import_folder_dict("graphics", "level", "bg"),
        "items": import_sub_folders("graphics", "items"),
        "player": import_sub_folders("graphics", "player")
    }

    self.level_frames.update(import_sub_folders("graphics", "level", "palms"))

    self.font = pygame.font.Font(join("graphics", "ui", "runescape_uf.ttf"), 40)
    self.ui_frames = {
        "heart": import_folder("graphics", "ui", "heart"),
        "coin": import_image("graphics", "ui", "coin")
    }

    self.overworld_frames = {
        "palms": import_folder("graphics", "overworld", "palm"),
        "water": import_folder("graphics", "overworld", "water"),

        # imported as nested dict
        "icon": import_sub_folders("graphics", "overworld", "icon"),
        "path": import_folder_dict("graphics", "overworld", "path")
    }

    self.audio = import_sound("audio")
