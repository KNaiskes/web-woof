# web-woof

A simple multithreaded watchdog for Websites.

web-woof gets and prints Website's status code.

[Useful - HTTP response status codes by Mozilla](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)

## Requirements

* Go >= 1.14

## Quick Start

Clone project

```
$ git clone https://github.com/KNaiskes/web-woof
```

### Append your Websites urls in the list

Open **web-woof.go** with your favorite editor and append the urls list
(**line: 12**)

```go
// Add comma separated your website's urls
// Note even the last one in the list must include a comma in the end
var urls = []string {
	// "https://www.example.com",
	// "https://www.example2.com",
}

```

### Run without compiling

```
$ cd web-woof
$ go run web-woof
```

### Compile it and then run it

```
$ cd web-woof
$ go build
$ ./web-woof
```
