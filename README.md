# HTTP Fetch for Tinh Tinh

<div align="center">
<img alt="GitHub Release" src="https://img.shields.io/github/v/release/tinh-tinh/fetch">
<img alt="GitHub License" src="https://img.shields.io/github/license/tinh-tinh/fetch">
<a href="https://codecov.io/gh/tinh-tinh/fetch" > 
 <img src="https://codecov.io/gh/tinh-tinh/fetch/graph/badge.svg?token=VK57E807N2"/> 
 </a>
<a href="https://pkg.go.dev/github.com/tinh-tinh/fetch"><img src="https://pkg.go.dev/badge/github.com/tinh-tinh/fetch.svg" alt="Go Reference"></a>
</div>

<div align="center">
    <img src="https://avatars.githubusercontent.com/u/178628733?s=400&u=2a8230486a43595a03a6f9f204e54a0046ce0cc4&v=4" width="200" alt="Tinh Tinh Logo">
</div>

## Overview

A flexible, type-safe HTTP client for the Tinh Tinh Framework, inspired by fetch/axios, supporting full-featured requests, easy configuration, and integration with Tinh Tinh modules.

## Install

```bash
go get -u github.com/tinh-tinh/fetch/v2
```

## Features

- Simple and chainable syntax for all HTTP verbs (GET, POST, PUT, PATCH, DELETE, etc.)
- Base URL and per-request configuration
- Automatic JSON encoding/decoding (customizable)
- Query parameter and header support
- Timeout and cancellation (context-based)
- Cookie and credential management
- Response formatting with type inference
- Full integration as a Tinh Tinh module/provider

## Quick Usage

### Create a Fetch Instance

```go
import "github.com/tinh-tinh/fetch/v2"

client := fetch.Create(&fetch.Config{
    BaseUrl: "https://jsonplaceholder.typicode.com",
    Headers: map[string][]string{
        "x-api-key": {"abcd", "efgh"},
    },
    Timeout: 5 * time.Second,
    ResponseType: "json",
})
```

### Make HTTP Requests

```go
// GET with query params
type Query struct {
    PostID int `query:"postId"`
}
res := client.Get("comments", &Query{PostID: 1})

// POST with body and query
type Post struct {
    UserID int    `json:"userId"`
    Title  string `json:"title"`
    Body   string `json:"body"`
}
res = client.Post("posts", &Post{UserID: 1, Title: "foo", Body: "bar"}, &Query{PostID: 1})

// PUT, PATCH, DELETE
client.Put("posts/1", &Post{UserID: 1, Title: "updated", Body: "new"}, &Query{PostID: 1})
client.Patch("posts/1", &Post{UserID: 1, Title: "patch", Body: "body"}, &Query{PostID: 1})
client.Delete("posts/1", &Query{PostID: 1})
```

## Contributing

We welcome contributions! Please feel free to submit a Pull Request.

## Support

If you encounter any issues or need help, you can:
- Open an issue in the GitHub repository
- Check our documentation
- Join our community discussions
