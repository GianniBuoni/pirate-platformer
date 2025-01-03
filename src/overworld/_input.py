from settings import *

def input(self):
    keys = pygame.key.get_pressed()
    valid_inputs = self.availabe_inputs()[0]

    if self.can_input:
        if keys[pygame.K_LEFT] and "left" in valid_inputs:
            self.start_point = self.icon.rect.topleft
            self.current_path = self.availabe_paths()["left"][1:]
            self.can_input = False
        if keys[pygame.K_RIGHT] and "right" in valid_inputs:
            self.start_point = self.icon.rect.topleft
            self.current_path = self.availabe_paths()["right"][1:]
            self.can_input = False
        if keys[pygame.K_UP] and "up" in valid_inputs:
            self.start_point = self.icon.rect.topleft
            self.current_path = self.availabe_paths()["up"][1:]
            self.can_input = False
        if keys[pygame.K_DOWN] and "down" in valid_inputs:
            self.start_point = self.icon.rect.topleft
            self.current_path = self.availabe_paths()["down"][1:]
            self.can_input = False
