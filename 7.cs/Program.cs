int sumOfDistancesPart1(List<int> positions, int index) {
  return positions.Aggregate(0, (sum, value) => sum + Math.Abs(value - index));
}

int sumOfDistancesPart2(List<int> positions, int index) {
  return positions.Aggregate(0, (sum, value) => {
    int n = Math.Abs(value - index);
    return sum + n * (n + 1) / 2;
  });
}

int result(List<int> positions, Func<List<int>, int, int> sumOfDistances) {
  int minSum = sumOfDistances(positions, 0);

  for (int i = 1; i < positions.Max(); i++) {
    int newSum = sumOfDistances(positions, i);
    if (newSum < minSum) minSum = newSum;
  }

  return minSum;
}

List<int> positions = new List<int>();
string[] lines = File.ReadAllLines("input.txt");  
positions = lines[0].Split(",").Select(Int32.Parse).ToList();

Console.WriteLine("Part 1: " + result(positions, sumOfDistancesPart1));
Console.WriteLine("Part 2: " + result(positions, sumOfDistancesPart2));
