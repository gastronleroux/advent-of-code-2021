from numpy import genfromtxt
arr = genfromtxt('input.txt', dtype='str')
radical_b = lambda arr, r, most=True: [((int(sum([int(arr[j][i]) for j in range(len(arr))]) >= len(arr) / 2)) + int(most)) % 2 for i in range(*r)]
bin_to_dec = lambda arr: int(''.join(map(str, arr)), 2)

# 3.1
print(bin_to_dec(radical_b(arr, [len(arr[0])])) * bin_to_dec(radical_b(arr, [len(arr[0])], most=False)))

# 3.2
def radical_b_extra(arr, r, most=True):
  for i in range(*r):
    if len(arr) == 1: break
    arr = [n for n in arr if n[i] == str(radical_b(arr, [i, i + 1], most=most)[0])]
  return arr[0]

print(bin_to_dec(radical_b_extra(arr, [len(arr[0])])) * bin_to_dec(radical_b_extra(arr, [len(arr[0])], most=False)))
