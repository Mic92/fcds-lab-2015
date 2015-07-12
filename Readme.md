My implementation Foundations of Concurrent and Distributed Systems Lab in go

Implemented algorithm:

- [x] bucketsort
- [ ] friendly
- [x] haar
- [ ] knapsack
- [x] threesat

## Usage:

```
$ cd fcds-lab-2015
$ go build .
```

### bucketsort

```
$ ./fcds-lab-2015 bucketsort < bucketsort/input/large.in > sorted
```

### threesat

```
$ ./fcds-lab-2015 threesat < bucketsort/input/large.in
```

### haar

```
$ gcc ./haar/input_generator.c -o ./haar/input_generator.o
$ ./haar/input_generator haar/input/large.in
$ ./fcds-lab-2015 haar < haar/input/large.in
```
