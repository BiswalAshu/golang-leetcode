def ret_string(given):
    count = -1
    new_string = list(given)

    for i in given:
        if i != '#':
            count += 1
            new_string[count] = i
        else:
            count -= 1

        if count < -1:
            count = -1

    return ''.join(new_string[:count+1])


def backspace_compare(s, t):
    if ret_string(s) == ret_string(t):
        return True
    return False
