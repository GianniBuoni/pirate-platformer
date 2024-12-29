class Data():
    def __init__(self) -> None:
        self._health = 5
        self._coins = 0

    @property
    def health(self):
        return self._health

    @health.setter
    def health(self, value):
        self._health = min(5, max(0, value))

    @property
    def coins(self):
        return self._coins

    @coins.setter
    def coins(self, value):
        self._coins = max(0, value)
