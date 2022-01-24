const fs = require("fs");
const arr = fs.readFileSync("./input.txt", "utf-8").split("\n").slice(0, -1).map(s => s.split(" ")).map(arr => [arr[0], parseInt(arr[1])]);

// 2.1
const dir = n => ({forward: [n[1], 0], up: [0, -n[1]], down: [0, n[1]]}[n[0]]);
console.log(arr.reduce((s, n) => dir(n).map((p, i) => p + s[i]), [0, 0]).reduce((s, n) => s * n, 1));

// 2.2
const dir = (n, aim) => ({forward: [0, n[1], n[1] * aim], up: [-n[1], 0, 0], down: [n[1], 0, 0]}[n[0]]);
console.log(arr.reduce((s, n) => dir(n, s[0]).map((p, i) => p + s[i]), [0, 0, 0]).slice(1).reduce((s, n) => s * n, 1));
