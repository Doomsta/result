# result

## Example
### Option
```golang
// TODO
```

### Result
```golang
func foo(in string) Result[string] {
    if in != "world" {
        return Failed[string](errors.New("in != world"))
    }
    return Ok[string]("hello, world")
}

func main() {
	switch v := foo("world").(type) {
	case Failed[string]:
		fmt.Println(v.Error())
	case OK[string]:
		fmt.Println(v.Unwrap())
	}

	if v := foo("world"); v.IsOk() {
		fmt.Println(v.Unwrap())
	}
}
```

### Nullable
```golang
// TODO
```

## JSON
Due to golang way to handle interface, you can't use json.Marshal directly.
There is no problem to implement MarshalJSON, but you can't unmarshal it back.
This would require to implement UnmarshalJSON, on an interface, which is not possible.

```golang
func (o *Option[T]) UnmarshalJSON(data []byte) error {
//      ^^^^^^^^^^ this is not possible
    var v T
    if err := json.Unmarshal(data, &v); err != nil {
        *o = None[T]()
        return nil
    }
    *o = Some[T](v)
    return nil
}
```
