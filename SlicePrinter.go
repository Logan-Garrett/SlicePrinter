package SlicePrinter

import (
    "fmt"
)

func SlicePrint(slice []string) {
    for _, element := range slice {
        fmt.Println(element)
    }
}
