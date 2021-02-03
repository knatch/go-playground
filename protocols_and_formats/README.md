# Protocols and Formats 

## Protocol Packages

"net/http" - web communication protocol
```golang
http.Get("knatch.me")
```

"net" - TCP/IP and socket programming
```golang
net.Dial("tcp", "knatch.me:80")
```

## JSON
- JavaScript Object Notation - RFC 7159
- Format to represent structured information
- Atribute-Value pairs -> `struct` or `map`

### Golang Object - Struct
```golang
p1 := Person(name: "Jane", addr: "123 A. Street", phone: "xx34")
// JSON
// {"name": "Jane", "addr": "123 A. Street", "phone": "xx34"}
```

### JSON Marshalling/Unmarshalling
`Marshal()` returns JSON representation as []byte

`Unmarshal()` converts a JSON []byte into a Go object
```golang
p1 := Person(name: "Jane", addr: "123 A. Street", phone: "xx34")

bArr, err := json.Marshal(p1)

var p2 Person
err := json.Unmarshal(bArr, &p2)
```


## File Access
Linear access, not random access
- Open - get handle for access
- Read - read bytes into []byte
- Write - write []byte into file
- Close - release handle
- Seek - move read/write head

### ioutil File Read
"io/ioutil" package
```golang
dat, e := ioutil.ReadFile("test.txt")
// dat is []byte filled with the content of entire file
// Explicit open/close is not needed
```
Large files can cause a problem

### ioutil File Write
```golang
dat = "hello world"

err := ioutil.WriteFile("out.txt", dat, 0777)
```

## os Package File Access
`os.Open` opens a file, returns a file descriptor

`os.Close` closes a file

`os.Read` reads from a file into []byte

`os.Write` writes a []byte into file

### os File Read
```golang
f, err := os.Open("test.txt")
bArr := make([]byte, 10)
nb, err := f.Read(bArr)
f.Close()
```

### os File Write
```golang
f, err := os.Create("out.txt")

bArr := []byte{1, 2, 3}
nb, err := f.Write(bArr)
// or
nb, err := f.WriteString("Hello")
```