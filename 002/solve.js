var temp, sum = 0, a = 0, b = 1

while (b < 4e6) {
  temp = a;
  a = b;
  b += temp;
  if ((b & 1) === 0) {
    sum += b;
  }
}

console.log(sum);

