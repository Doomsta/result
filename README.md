# result

## Example

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
