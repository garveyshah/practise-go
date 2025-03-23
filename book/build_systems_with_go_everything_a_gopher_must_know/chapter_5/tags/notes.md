## 5.4 TAGS

Go provides powerful and versatile structs enrichment using tags. These tags can be interpreted  on run-time using reflection which adds valuable information that can be employed for different purposes.

```go

type User struct {
    userId string `tagA: "valueA1" tagB: "valueA2"
    Email string `tagB: "value"
    Password string `tagC: "v1 v2"`
    Others string `"something a b"`
}
```
The `struct` above decalares four fields with different tags. Every tag becomes an attribute of the field that can be accessed later. Go permits tags to declare raw string literals like `"something a b"`. However, by convection tags follow a key-value schema separated by spaces. For example, the string `tagA: "valueA1" tagB:"valueA2"`, decalres tow tags `tagA` and `tagB` with values `valueA1` and `valueA2` respectively.
