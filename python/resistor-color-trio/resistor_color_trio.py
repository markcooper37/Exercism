def label(colors):
    color_values = ["black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"]
    number = (10 * color_values.index(colors[0]) + color_values.index(colors[1])) * 10 ** (color_values.index(colors[2]))
    if number < 1000:
        return str(number) + " ohms"
    if number < 1000000:
        return str(number // 1000) + " kiloohms"
    if number < 1000000000:
        return str(number // 1000000) + " megaohms"
    return str(number // 1000000000) + " gigaohms"
    
