/*

Link: https://neetcode.io/problems/time-based-key-value-store/question?list=neetcode150

Design a time-based key-value data structure that can store multiple values for the same key at different time stamps and retrieve the key's value at a certain timestamp.

Implement the TimeMap class:

TimeMap() Initializes the object of the data structure.
void set(String key, String value, int timestamp) Stores the key key with the value value at the given time timestamp.
String get(String key, int timestamp) Returns a value such that set was called previously, with timestamp_prev <= timestamp. If there are multiple such values, it returns the value associated with the largest timestamp_prev. If there are no values, it returns "".
Example 1:

Input:
["TimeMap", "set", ["alice", "happy", 1], "get", ["alice", 1], "get", ["alice", 2], "set", ["alice", "sad", 3], "get", ["alice", 3]]

Output:
[null, null, "happy", "happy", null, "sad"]

Explanation:
TimeMap timeMap = new TimeMap();
timeMap.set("alice", "happy", 1);  // store the key "alice" and value "happy" along with timestamp = 1.
timeMap.get("alice", 1);           // return "happy"
timeMap.get("alice", 2);           // return "happy", there is no value stored for timestamp 2, thus we return the value at timestamp 1.
timeMap.set("alice", "sad", 3);    // store the key "alice" and value "sad" along with timestamp = 3.
timeMap.get("alice", 3);           // return "sad"
Constraints:

1 <= key.length, value.length <= 100
key and value only include lowercase English letters and digits.
0 <= timestamp <= 10^7
All the timestamps of set are strictly increasing.
*/

package main

import (
	"fmt"
	"sort"
)

type TimeMap struct {
	m map[string][]pair
}

type pair struct {
	timestamp int
	value     string
}

func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][]pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], pair{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, exists := this.m[key]; !exists {
		return ""
	}

	pairs := this.m[key]
	idx := sort.Search(len(pairs), func(i int) bool {
		return pairs[i].timestamp > timestamp
	})

	if idx == 0 {
		return ""
	}
	return pairs[idx-1].value
}

func main() {
	timeMap := Constructor()
	timeMap.Set("alice", "happy", 1)     // store the key "alice" and value "happy" along with timestamp = 1.
	fmt.Println(timeMap.Get("alice", 1)) // return "happy"
	fmt.Println(timeMap.Get("alice", 2)) // return "happy", there is no value stored for timestamp 2, thus we return the value at timestamp 1.
	timeMap.Set("alice", "sad", 3)       // store the key "alice" and value "sad" along with timestamp = 3.
	fmt.Println(timeMap.Get("alice", 3)) // return "sad"
}
