# HAM (Highly Accountable Math)

HAM is a go package that provides an interface for performing high precision math with a built in audit log showing the operations that occurred to achieve an end result.

```go
ctx := ham.WithContext(ctx)
exponent := ham.IntConstant(ctx, 3)
input := ham.IntVariable(ctx, userInput)
result := input.Pow(ctx, exponent)
```

That will produce the following equation

```
a = const 3
b = var 15
c = (b^a) 3375
```

Variables can also be named to make the results better

```go
ctx := ham.WithContext(ctx)
exponent := ham.IntConstant(ctx, 3, ham.VarName("Exponent"))
input := ham.IntVariable(ctx, userInput, ham.VarName("Input"))
result := input.Pow(ctx, exponent, ham.VarName("Result"))
```

```
Exponent = const 3
Input = var 15
Result = (Input^Exponent) 3375
```