# URLMatcher GoLang

URLMatcher useful for url matching.

```
func main() {
    m := NewMatcher("/users/{id:[0-9]*}")

    match, err := m.Match("/users/1")

    if match.Has("id") {
        id := match.Var("id")
        fmt.Println(id)
    }
}
```