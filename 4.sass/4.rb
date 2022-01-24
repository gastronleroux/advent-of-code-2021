lines = File.readlines("input.txt").filter{|line| line != "\n" }.map(&:chomp)

numbers = "$numbers: (#{lines[0]})"
boards = "$boards: ("
(0..(lines.length - 1)/5 - 1).each do |i|
  line = lines[i * 5 + 1..i * 5 + 5].map do |line|
    "(" + line.gsub(/\s+/m, ' ').strip.split(" ").join(",") + ")"
  end
  boards += "(" + line.join(",") + "),"
end
boards += ")"
content = numbers + "\n" + boards

File.open("4_helper.sass", "w") {|f| f.write(content) }
