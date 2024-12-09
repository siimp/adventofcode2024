numbers = []

def is_safe(numbers):
    max = 3
    is_increasing = numbers[0] < numbers[1]
    for i in range(0, len(numbers) - 1):
        if is_increasing and numbers[i] >= numbers[i + 1]:
            return False
        if is_increasing and numbers[i + 1] - numbers[i] > max:
            return False
        if not is_increasing and numbers[i] <= numbers[i + 1]:
            return False
        if not is_increasing and numbers[i] - numbers[i + 1]  > max:
            return False

    return True

safe_reports = 0
with open("input.txt") as file:
    report = -1
    for line in file:
        numbers = [int(n) for n in line.split(" ")]
        safe = is_safe(numbers)
        if safe:
            safe_reports+=1
            continue

        # try each variant
        for i in range(0, len(numbers)):
            modified_numbers = numbers[0:i] + numbers[i+1:len(numbers)]
            safe = is_safe(modified_numbers)
            if safe:
                safe_reports+=1
                break

print(f'safe reports: {safe_reports}')