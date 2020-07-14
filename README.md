# egu
easy golang utils

## import
```
go get github.com/ebar-go/egu
```
## String
- UUID: generate unique id
```
id := egu.UUID()
```

- DefaultString: return default value when string is empty
```
var s string
s = egu.DefaultString(s, "defaultValue")
```

## Number
