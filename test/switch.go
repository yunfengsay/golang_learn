package main
import (
    "encoding/json"
    "fmt"
    "reflect"
)
var txt = `
[
[
    "define",
    "a",
    [
    "read",
    "cd"
    ],
    [
    "if",
    [
        "\u003e",
        "a",
        "cd"
    ],
    [
        "print",
        "demo"
    ],
    [
        "print",
        "not demo"
    ]
    ]
],
[
    "say",
    [
    "print",
    "a",
    [
        "save",
        [
        "b",
        [
            "x",
            [
            "c",
            "8"
            ]
        ]
        ]
    ]
    ]
],
[
    "print",
    "fun"
]
]`
func Switch(value reflect.Value) {
    kind := value.Kind()
    switch kind {
    case reflect.String:
        fmt.Printf("string: %s\n", value.String())
    case reflect.Struct:
        parseStruct(value)
    case reflect.Slice:
        parseSlice(value)
    case reflect.Interface:
        v := reflect.ValueOf(value.Interface())
        Switch(v)
    default:
        fmt.Printf("unknown type: %s\n", kind.String())
    }
}
func parseSlice(value reflect.Value) {
    size := value.Len()
    for j := 0; j < size; j++ {
        field := value.Index(j)
        Switch(field)
    }
}
func parseStruct(value reflect.Value) {
    size := value.NumField()
    for i := 0; i < size; i++ {
        field := value.Field(i)
        Switch(field)
    }
}
func main() {
    var objs []interface{}
    if err := json.Unmarshal([]byte(txt), &objs); err != nil {
        fmt.Println(err)
        return
    }
    for _, obj := range objs {
        Switch(reflect.ValueOf(obj))
    }
}
