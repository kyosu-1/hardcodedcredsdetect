# hardcodedcredsdetect

`hardcodedcredsdetect` is a static analysis tool for Go that detects hard-coded secrets in source code and helps identify security risks.

## Features

`hardcodedcredsdetect` provides the following features:

- Detect if variables that may contain sensitive information are hard-coded

## Installation

You can install `hardcodedcredsdetect` using the following command:

```shell
$ go install github.com/kyosu-1/hardcodedcredsdetect
```

## Usage

You can use `hardcodedcredsdetect` by running the following command:

```shell
$ go vet -vettool=which hardcodedcredsdetect <package_directory>
```

## Variable Name Pattern
hardcodedcredsdetect uses the following regular expression to detect variable names that may contain sensitive information:

* `(?i)password|passwd|pwd`
* `(?i)credential|cred|auth.*token|api.*key`

## Examples

The following is an example of code that `hardcodedcredsdetect` will detect:

```go
package main

import "fmt"

func main() {
    password := "password123" // This will be detected by hardcodedcredsdetect
    fmt.Println(password)
}
```
