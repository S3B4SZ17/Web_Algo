import sys

num = int(sys.argv[1])
reversed_num = 0

while num != 0:
    digit = num % 10
    reversed_num = reversed_num * 10 + digit
    num //= 10

print(str(reversed_num))
