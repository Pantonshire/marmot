package marmot

import (
    "fmt"
    "strings"
    "testing"
)

func TestHTML(t *testing.T) {
    expect := strings.TrimSpace(`
<html lang="en">
<head>
    <title>Test</title>
</head>
<body>
    <h1>Hello world</h1>
    <p>Foo</p>
    <p>Baa</p>
</body>
</html>`)

    var cache HTMLCache
    if err := cache.Load(Directory{Path: "testdata/html", Extension: ".gohtml"}); err != nil {
        t.Error(err)
    }

    str, err := cache.Builder("Page").ExecStr()
    if err != nil {
        t.Error(err)
    }

    fmt.Println(str)

    if strings.TrimSpace(str) != expect {
        t.Error("incorrect html output")
    }
}
