arr = File.readlines("input.txt").map(&:to_i)

# 1.1
puts arr[1..arr.length].reduce([0, arr[0]]) {|s, n| [s[0] + (n > s[1] ? 1 : 0), n]}[0]

# 1.2
puts (2...arr.length).reduce([0, arr[0..2].sum]) {|s, n| [s[0] + (arr[n - 2..n].sum > s[1] ? 1 : 0), arr[n - 2..n].sum]}[0]
