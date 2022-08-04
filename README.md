# url

A [Golang](http://golang.org/) package for URL parsing and normalization.

1. [Parsing URL](#parsing-url)
2. [Normalizing URL](#normalizing-url)

## Parsing URL

### Difference between `github.com/hueristiq/url` and `net/url`

<table>
<thead>
<tr>
<th><a href="https://godoc.org/github.com/hueristiq/url#Parse">github.com/hueristiq/url</a></th>
<th><a href="https://golang.org/pkg/net/url/#Parse">net/url</a></th>
</tr>
</thead>
<tr>
<td>
<pre>
url.Parse("example.com")

&url.URL{
   Scheme:  "http",
   Host:    "example.com",
   Path:    "",
}
</pre>
</td>
<td>
<pre>
url.Parse("example.com")

&url.URL{
   Scheme:  "",
   Host:    "",
   Path:    "example.com",
}
</pre>
</td>
</tr>
<tr>
<td>
<pre>
url.Parse("localhost:8080")

&url.URL{
   Scheme:  "http",
   Host:    "localhost:8080",
   Path:    "",
   Opaque:  "",
}
</pre>
</td>
<td>
<pre>
url.Parse("localhost:8080")

&url.URL{
   Scheme:  "localhost",
   Host:    "",
   Path:    "",
   Opaque:  "8080",
}
</pre>
</td>
</tr>
<tr>
<td>
<pre>
url.Parse("user.local:8000/path")

&url.URL{
   Scheme:  "http",
   Host:    "user.local:8000",
   Path:    "/path",
   Opaque:  "",
}
</pre>
</td>
<td>
<pre>
url.Parse("user.local:8000/path")

&url.URL{
   Scheme:  "user.local",
   Host:    "",
   Path:    "",
   Opaque:  "8000/path",
}
</pre>
</td>
</tr>
</table>

### Usage

```go
import "github.com/hueristiq/url"

func main() {
    url, _ := url.Parse("example.com")
    // url.Scheme == "http"
    // url.Host == "example.com"

    fmt.Print(url)
    // Prints http://example.com
}
```

## Normalizing URL

The [url.Normalize()](https://godoc.org/github.com/hueristiq/url#Normalize) function normalizes the URL using the predefined subset of [Purell](https://github.com/PuerkitoBio/purell) flags.

### Usage

```go
import "github.com/hueristiq/url"

func main() {
    url, _ := url.Parse("localhost:80///x///y/z/../././index.html?b=y&a=x#t=20")
    normalized, _ := url.Normalize(url)

    fmt.Print(normalized)
    // Prints http://localhost/x/y/index.html?a=x&b=y#t=20
}
```