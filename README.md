# Go Badge

[![Tests & Linting](https://github.com/hatamiarash7/gobadge/actions/workflows/test.yml/badge.svg)](https://github.com/hatamiarash7/gobadge/actions/workflows/test.yml)

GoBadge is an open source project for displaying information via SVG "badges". GoBadge was inspired by the open source project, **[shields.io](https://shields.io)**.

## Getting Started

```bash
go get github.com/hatamiarash7/gobadge
```

```go
package main

import (
    "github.com/hatamiarash7/gobadge"
)
```

## Example

Run server:

```bash
cd example && go run server.go
```

Then get your badge from `http://localhost:8080/example?label=hello&tag=world&color=green`

## Known problem

Currently, badge will not be rendered correctly with some letters in the label because of formula for dynamic width calculation.

![example 1](.github/exp1.png)

![example 1](.github/exp2.png)

The length of Capital letters is bigger than other ones. ( I will working on it )
