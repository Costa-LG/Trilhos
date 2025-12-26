The & Operator

& goes in front of a variable when you want to get that variable's memory address.

The * Operator

* goes in front of a variable that holds a memory address and resolves it (it is therefore the counterpart to the & operator). It goes and gets the thing that the pointer was pointing at, e.g. *myString.

```go
myString := "Hi"
fmt.Println(*&myString)  // prints "Hi"
```
or more usefully, something like

```go
myStructPointer = &myStruct
// ...
(*myStructPointer).someAttribute = "New Value"
```

* in front of a Type

When * is put in front of a type, e.g. *string, it becomes part of the type declaration, so you can say "this variable holds a pointer to a string". For example:

var str_pointer *string
So the confusing thing is that the * really gets used for 2 separate (albeit related) things. The star can be an operator or part of a type.