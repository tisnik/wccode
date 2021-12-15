from collections import Counter

counter = Counter()

with open("results.txt") as fin:
    for i, line in enumerate(fin):
        type = line.split(" ")[1][:-1]
        counter[type] += 1

for cnt, type in counter.most_common(30):
    print(cnt, type)
