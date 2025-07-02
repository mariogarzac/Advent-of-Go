from collections import Counter

def read_file():
    left = []
    right = []

    with open("input.txt") as file:
        for line in file:
            numbers = (line.strip()).split()

            left.append(int(numbers[0]))
            right.append(int(numbers[1]))

    return left, right

def part1(left, right):
    # dif = 0
    # for i in range(len(left)):
    #     dif += abs(left[i]-right[i])
    #
    # return dif
   return sum(abs(l - r) for l, r in zip(left, right))

# def part2(left, right):
#     num_map = dict()
#
#     for i in range(len(left)):
#         count = 0
#         if left[i] in num_map:
#             continue
#         for j in range(len(right)):
#             if left[i] == right[j]:
#                 count += 1
#
#         num_map[left[i]] = count
#
#     total = 0
#     for num in left:
#         total += (num * num_map[num])
#
#     return total

def part2(left, right):
    # Create a frequency map for the right list
    right_count = Counter(right)

    # Calculate the total based on the frequency map
    total = sum(num * right_count[num] for num in left)
    return total

def main():
    left, right = read_file()
    left.sort()
    right.sort()

    print(part1(left, right))
    print(part2(left, right))

main()
