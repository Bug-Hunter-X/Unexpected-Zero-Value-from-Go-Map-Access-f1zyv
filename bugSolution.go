```Go
func main() {
    var m = make(map[string]int)
    m["a"] = 1

    value, ok := m["b"]
    if ok {
        fmt.Println(value) // Access value only if key exists
    } else {
        fmt.Println("Key 'b' not found")
    }

    value, ok = m["c"]
    if ok {
        fmt.Println(value) // Access value only if key exists
    } else {
        fmt.Println("Key 'c' not found")
    }
}
```