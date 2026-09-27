# Markdown Code Block Runner Matrix

The paired blocks exercise isolated file runs and persistent REPL state. For a REPL-capable language, run the first
block once, then run the second block twice: the visible values must be `1`, `2`, and `3`. File mode runs each block in
a fresh temporary directory, so the first block prints `1` and the second block must fail because `a` is undefined.

Go and Rust are file-mode only. Java uses separate file and JShell pairs because declaring a `Main` class in JShell does
not execute its `main` method.

## Python

```python

a = 1
print(a)

```

```python
a = a + 1

print(a)
```

## JavaScript

```js
var a = 1;
console.log(a);
```

```js
a = a + 1;
console.log(a);
```

## Bash

```bash
a=1
printf '%s\n' "$a"
```

```bash
if [[ ! -v a ]]; then
  printf '%s\n' 'a is undefined' >&2
  false
else
  ((a += 1))
  printf '%s\n' "$a"
fi
```

## Go — file mode only

```go
package main

import "fmt"

func main() {
    a := 1
    fmt.Println(a)
}
```

```go
package main

import "fmt"

func main() {
    a = a + 1
    fmt.Println(a)
}
```

## Elixir

```elixir
a = 1
IO.puts(a)
```

```elixir
a = a + 1
IO.puts(a)
```

## F\#

```fsharp
let mutable a = 1
printfn "%d" a
```

```fsharp
a <- a + 1
printfn "%d" a
```

## C\#

```csharp
var a = 1;
Console.WriteLine(a);
```

```csharp
a = a + 1;
Console.WriteLine(a);
```

## Java — file mode

```java
public class Main {
    public static void main(String[] args) {
        int a = 1;
        System.out.println(a);
    }
}
```

```java
public class Main {
    public static void main(String[] args) {
        a = a + 1;
        System.out.println(a);
    }
}
```

## Java — JShell

```java
int a = 1;
System.out.println(a);
```

```java
a = a + 1;
System.out.println(a);
```

## Lua

```lua
a = 1
print(a)
```

```lua
a = a + 1
print(a)
```

## Rust — file mode only

```rust
fn main() {
    let a = 1;
    println!("{a}");
}
```

```rust
fn main() {
    a += 1;
    println!("{a}");
}
```
