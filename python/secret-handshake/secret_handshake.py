def commands(binary_str):
    handshake = [action for index, action in enumerate(['wink', 'double blink', 'close your eyes', 'jump']) if binary_str[4 - index] == '1']
    if binary_str[0] == '1':
        handshake.reverse()
    return handshake
