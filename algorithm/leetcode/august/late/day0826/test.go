package main

import (
	"fmt"
	"strings"
)

func main() {
	str := `
[["Lily","Lily4@m.co","Lily5@m.co"],["Lily","Lily8@m.co","Lily9@m.co"],["Lily","Lily15@m.co","Lily16@m.co"],["Lily","Lily19@m.co","Lily20@m.co"],["Lily","Lily6@m.co","Lily7@m.co"],["Lily","Lily10@m.co","Lily11@m.co"],["Lily","Lily5@m.co","Lily6@m.co"],["Lily","Lily13@m.co","Lily14@m.co"],["Lily","Lily9@m.co","Lily10@m.co"],["Lily","Lily1@m.co","Lily2@m.co"],["Lily","Lily3@m.co","Lily4@m.co"],["Lily","Lily2@m.co","Lily3@m.co"],["Lily","Lily11@m.co","Lily12@m.co"],["Lily","Lily7@m.co","Lily8@m.co"],["Lily","Lily12@m.co","Lily13@m.co"],["Lily","Lily18@m.co","Lily19@m.co"],["Lily","Lily17@m.co","Lily18@m.co"],["Lily","Lily16@m.co","Lily17@m.co"],["Lily","Lily14@m.co","Lily15@m.co"],["Lily","Lily0@m.co","Lily1@m.co"]]`
	str = strings.ReplaceAll(str, "[", "{")
	str = strings.ReplaceAll(str, "]", "}")
	fmt.Println(str)
}
