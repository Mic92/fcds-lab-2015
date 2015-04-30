package bucketsort

import (
	"fmt"
)

const (
	NUMBER_OF_BUCKETS = 94
)

func Sort(data []byte, length int) []int {
	var buckets [NUMBER_OF_BUCKETS][]int
	size := len(data)
	var idx byte

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%d\n", idx)
		}
	}()

	returns := make([]int, len(data))
	for i := range buckets {
		offset := (i * size) / NUMBER_OF_BUCKETS
		buckets[i] = returns[offset:offset:len(returns)]
	}

	elements := size / length
	for i := 0; i < elements; i++ {
		key := data[i*length] - 0x21
		//fmt.Printf("append(buckets[%d], %d (data[%d]=%s))\n", key, i, i*length, data[i*length:i*length+length])
		//fmt.Printf("returns=%v\n", returns)
		//fmt.Printf("buckets[%d] -> %v\n", key, buckets[key])
		//idx = data[i*length]
		fmt.Println(data[i*length : i*length+10])
		idx = key
		buckets[key] = append(buckets[key], i)
		//fmt.Printf("buckets[%d] -> %v\n", key, buckets[key])
	}

	return returns
}

//#include <stdlib.h>
//#include <string.h>
//#include "bucketsort.h"
//
//#define N_BUCKETS 94
//
//typedef struct {
//	long int *data;
//	int length;
//	long int total;
//} bucket;
//
//void sort(char *a, bucket *bucket) {
//	int j, i, length;
//	long int key;
//	length = bucket->length;
//	for (j = 1; j < bucket->total; j++) {
//		key = bucket->data[j];
//		i = j - 1;
//		while (i >= 0
//				&& strcmp(a + bucket->data[i] * length, a + key * length) > 0) {
//			bucket->data[i + 1] = bucket->data[i];
//			i--;
//		}
//		bucket->data[i + 1] = key;
//	}
//}
//
//long int* bucket_sort(char *a, int length, long int size) {
//
//	long int i;
//	bucket buckets[N_BUCKETS], *b;
//	long int *returns;
//
//	// allocate memory
//	returns = malloc(sizeof(long int) * size);
//	for (i = 0; i < N_BUCKETS; i++) {
//		buckets[i].data = returns + i * size / N_BUCKETS;
//		buckets[i].length = length;
//		buckets[i].total = 0;
//	}
//
//	// copy the keys to "buckets"
//	for (i = 0; i < size; i++) {
//		b = &buckets[*(a + i * length) - 0x21];
//		b->data[b->total++] = i;
//	}
//
//	// sort each "bucket"
//	for (i = 0; i < N_BUCKETS; i++)
//		sort(a, &buckets[i]);
//
//	return returns;
//}
