from enum import Enum

class CollidesWith(Enum):
    FLOOR = "floor"
    LEFT = "left"
    RIGHT = "right"
    AIR = "air"

class PlayerState(Enum):
    AIR_ATK = "air_attack"
    ATTACK = "attack"
    FALL = "fall"
    HIT = "hit"
    IDLE = "idle"
    JUMP = "jump"
    RUN = "run"
    WALL = "wall"
