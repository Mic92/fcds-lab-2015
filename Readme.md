My implementation Foundations of Concurrent and Distributed Systems Lab in go

Implemented algorithm:

- [x] bucketsort
- [ ] friendly
- [x] haar
- [ ] knapsack
- [x] threesat

## Usage:

```
$ git clone https://github.com/Mic92/fcds-lab-2015.git
$ cd fcds-lab-2015
$ git submodule update --init
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
$ cd lab/c_sequential/haar/ && make && bin/input_generator input/large.in && cd -
$ ./fcds-lab-2015 haar < lab/c_sequential/haar/input/large.in
```
