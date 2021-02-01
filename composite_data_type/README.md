# Composite Data Type

## Arrays
*fixed length* series of elements of a chosen type
Elements accesses using subscript notation []
Elements initialized to zero value

```golang
var x [5]int
x[0] = 2
// prints 2, 0, 0, 0, 0
fmt.Printf(x)
```

### Array Literal
An array with pre-defined values
```golang
// length of literal must equal length of array
var x [5]int = [5]{1, 2, 3, 4, 5}

// ... for size in array literal infers size from number of initializers
x := [...]{1, 2, 3, 4}
```

### Iterate Through Arrays
```golang
x := [3]int {1, 2, 3}

for i, v range x {
    fmt.Printf("ind %d, val %d", i, v)
}
```
## Slices
A "window" of an underlying array
Variable size, up to whole array

`Pointer` -  indicates start of the slice

`Length` - number of elements in the slice `len()`

`Capacity` - maximum number of elements from start to end of array `cap()`

```golang
arr := [...]string{"a", "b", "c", "d", "e", "f"}
s1 := arr[1:3]
s2 := arr[2:5]

// prints 2 6
fmt.Printf(len(s1), cap(s1))
```

### Acessing Slices
Writing to a slice changes underlying array
Overlapping slices refer to the same array elements

### Slice Literal
Can be used to initialize a slice
Creates the underlying array and creates a slice to reference it
Slice points to the start of the array, length is capacity

```golang
// you'd put the length if it was an array
sli := []int{1, 2, 3}
```
### Slice functions
make()
append()

## Hash Table
contains key/value pairs
used to compute the slot for a key

| Advantages | Disadvantages |
| ----------- | ----------- |
| faster lokup (constant-time vs linear-time) | may contain collisions (two keys hash to same slot) |
| arbitrary keys (not ints like slices or arrays) | |


### Maps
implementation of a hash table

```golang
var idMap map[string][int]
idMap = make(map[string]int)

// map literal
idMap := map[string]int {
    "jane": 1
}
```
#### Accessing Maps
referencing a value with [key]
```golang
// returns zero if key is not present
// adds a key/value pair
idMap["jano"] = 3

// deletes
delete(idMap, "jano")
```

#### Map functions
```golang
// id is value, p is presence of key (true/false)
id, p := idMap["joe"]

fmt.Println(len(idMap))

Iterating
// two values assignment with range
for key, val := range idMap {
    fmt.Printlmm(key, val)
}
```

## Struct
aggregate data type - groups together other objects of arbitrary type
```golang
type struct Person {
    name string
    addr string
    phone string
}

var p1 Person
// dot notation
p1.name = "Joe"
x = p1.addr

// initialized fields to zero with new
p1 := new(Person)

// struct literal
p2 := Person(name: "Jane", addr: "a street", phone: "3124")
```