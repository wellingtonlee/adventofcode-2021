nums = sorted([int(x) for x in open('input.txt', 'r').read().strip('\n').split(',')])

summorials = {}

for i in range(max(nums)+1):
    if i - 1 not in summorials:
        summorials[i] = 0
    else:
        summorials[i] = i + summorials[i-1]

min_num = -1
for i in range(max(nums)):
    tot = 0
    for j in nums:
        tot += summorials[abs(j-i)]
    if tot < min_num or min_num == -1:
        min_num = tot

print(f"Min: {min_num}")