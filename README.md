# result

## Example

```golang
func ExampleResultSwitch() {
	fn := func() Result[string] {
		return Ok[string]("hello")
	}

	switch v := fn().(type) {
	case Failed[string]:
		fmt.Println(v.Error())
	case OK[string]:
		fmt.Println(v.Unwrap())
	}

	if v := fn(); v.IsOk() {
		fmt.Println(v.Unwrap())
	}
}
```
