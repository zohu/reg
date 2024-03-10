### A regexp library that can be called in a chain.

```
func main() {
    // true
    reg.String("xx@xx.com").IsEmail().B()
    // true
    reg.Version("1.2.3").
        IsVersion().
        HighThan("1.2.2").
        LowThan("1.2.4").
        LowThan("1.*").
        Support("~1.2.0").
        Support("^1.2.0").
        Support(">=1.2.0").
        Support("<=2.2.0").
        B()
}
```