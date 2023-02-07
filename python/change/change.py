def find_fewest_coins(coins, target):
    if target < 0:
        raise ValueError("target can't be negative")
    if target == 0:
        return []
    change = []
    for i in range(len(coins) - 1, -1, -1):
        try:
            best = find_fewest_coins(coins[:i+1], target - coins[i])
            if len(best) < len(change)-1 or len(change) == 0:
                best.append(coins[i])
                change = best
        except:
            pass
        finally:
            if len(change) > 0 and i != 0 and target/coins[i-1] >= len(change):
                break
    if len(change) == 0:
        raise ValueError("can't make target with given coins")
    return change
