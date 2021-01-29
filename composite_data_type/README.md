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

## Variable Slices

## Hash Tables

## Maps

## Structs
