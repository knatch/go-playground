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

## Variable Slices

## Hash Tables

## Maps

## Structs
