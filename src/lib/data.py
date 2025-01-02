from typing import TYPE_CHECKING
if TYPE_CHECKING: from ui import UI

class Data():
    def __init__(self, ui: "UI") -> None:
        # player data
        self._health = 5
        self._coins = 0

        # level and stage data
        self.current_level = 5
        self.unlocked_levels = 5

        # ui init
        self.ui = ui
        self.ui.create_hearts(self.health)

    @property
    def health(self):
        return self._health

    @health.setter
    def health(self, value):
        self._health = min(5, max(0, value))
        self.ui.create_hearts(self._health)

    @property
    def coins(self):
        return self._coins

    @coins.setter
    def coins(self, value):
        self._coins = max(0, value)
        self.ui.show_coin(value)
