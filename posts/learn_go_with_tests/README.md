---
author: Jacob O'Neill
date: 20/10/2025
tags: ["books"]
urls: {"github": "https://github.com/quii/learn-go-with-tests", "gitbook": "https://quii.gitbook.io/learn-go-with-tests"}
---
# Learn Go with Tests

***Learn Go with Tests*** is a book about learning Go following the TDD philosophy *"Red, Green, Refactor"*. I though this would be a great book for myself to work through for my [ssgo](https://www.github.com/jacoboneill/ssgo) project as I wanted to practice my TDD and Go development with that project and get this blog online.

I found the book while scrolling through TikTok and YouTube and thought I would give it a go. This is going to be a lot longer of a post than I typically do. So here is the tl;dr, otherwise enjoy!

> TODO: Add tl;dr

## Introduction

The introduction does a lot of things for the reader, it explains how the writer has tried to teach developers Go using a myriad of techniques, none of them working quite as well as they would hope. On top of this, it explains how developers are reluctant to do TDD as it adds more work (my friend who works in QA has also mentioned this). They then go on to explain how to install Go (on MacOS it is as simple as `brew install go`). They then go on to explain the tools that Go comes with (`go mod` for setting up modules, `go fmt` for formatting code). We then go on to writing some code.

## Hello, World

The infamous *Hello, World*, a developers first steps into a new language, but this time with testing in mind. They split the program up to create a function that generates the `Hello, World!` `string`, and then pass that function into `main` to print to the screen. This `Hello` function can then be tested on to make sure that the returned string is in fact `Hello, World!`. Seems pretty redundant for a test until you read on and it explains about refactoring. It states that if you have unit tests for behaviour, then it is easier to catch out mistakes when refactoring. They give this excellent example of refactoring in the introduction:

```go
// Before refactoring
func Hello(name, language string) (string, error) {
    if language == "es" {
        return "Hola, " + name, nil
    } else if langauge == "fr" {
        return "Bonjour, " + name, nil
    } else if language == "en" {
        return "Hello, " + name, nil
    } else {
        return "", fmt.Errorf("Language not recognised: %q", language)
    }
}

// After refactoring
func Hello(name, language string) (string, error) {
    m := map[string]string{
        "en": "Hola",
        "fr": "Bonjour",
        "en": "Hello"
    }

    hello, ok := m[language]
    if !ok {
        return "", fmt.Errorf("Language not recognised: %q", language)
    }

    return fmt.Sprintf("%s, %s", hello, name), nil
}
```

The first maybe was created because only English was needed, then it got a feature request to use French and Spanish. What would've happened next is someone would've written tests for each of these cases, as well as the error. Then they decided to refactor using a `map` to make it easier to add more languages, as well as be able to consistently update the output without updating every language. The point is that they didn't change the behaviour when doing the refactor, so therefore the tests should still pass the same as they did before. If they don't then obviously something went wrong.

I finally wrote my [code](https://www.github.com/jacoboneill/blog/blob/main/posts/learn_go_with_tests/src/001_Hello_World/hello.go):
```go
package main

import "fmt"

func Hello() string {
	return "Hello, World!"
}

func main() {
	fmt.Println(Hello())
}
```

and [test](https://www.github.com/jacoboneill/blog/blob/main/posts/learn_go_with_tests/src/001_Hello_World/hello_test.go)
```go
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
```

When I do `go test` I get:
```
> go test
PASS
ok      001_Hello_World 0.342s
```
