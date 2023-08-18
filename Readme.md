# PathMatch

PathMatch is mostly used for routing.

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
