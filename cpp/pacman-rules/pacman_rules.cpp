// eat_ghost returns a boolean value if Pac-Man is able to eat the ghost.
bool can_eat_ghost(bool power_pellet_active, bool touching_ghost) {
    return power_pellet_active && touching_ghost;
}

// score returns a boolean value if Pac-Man scored.
bool scored(bool touching_power_pellet, bool touching_dot) {
    return touching_power_pellet || touching_dot;
}

// lost returns a boolean value if Pac-Man loses.
bool lost(bool power_pellet_active, bool touching_ghost) {
    return !power_pellet_active && touching_ghost;
}

// won returns a boolean value if Pac-Man wins.
bool won(bool has_eaten_all_dots, bool power_pellet_active, bool touching_ghost) {
    return has_eaten_all_dots && !lost(power_pellet_active, touching_ghost);
}