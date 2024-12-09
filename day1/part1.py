list1 = []
list2 = []
with open("input.txt") as file:
    for line in file:
        numbers = line.split("   ")
        list1.append(int(numbers[0]))
        list2.append(int(numbers[1]))


list1 = sorted(list1)
list2 = sorted(list2)

result = 0
for i, v in enumerate(list1):
    result += abs(list1[i] - list2[i])
print(result)