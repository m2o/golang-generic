package main

import (
    "fmt"
)

type Fetcher interface {
    // Fetch returns the body of URL and
    // a slice of URLs found on that page.
    Fetch(url string) (body string, urls []string, err error)
}

// struct renamed crawlShared to reflect that it provides access to
// things that are shared across goroutines (the goroutines being
// instances of c2, specifically.)
type crawlShared struct {
    Fetcher     // embedded for Fetch method.
 
    // the "salt shaker" pattern:  a channel holds the shared map
    mapAccess   chan map[string]bool
 
    // about the same pattern, except that there is nothing obvious
    // to put in the channel.  instead, a boolean value is used as
    // a token.  the value is ignored.  all that matters is the
    // presence or absence of the token in the channel.
    printAccess chan bool
}
 
func Crawl(url string, depth int, fetcher Fetcher) {

    c := &crawlShared{
        fetcher,
        make(chan map[string]bool, 1),
        make(chan bool, 1),
    }
 
    // put the salt shaker on the table.  that is, put the map
    // in the channel, making it available to goroutines.
    c.mapAccess <- map[string]bool{url: true}
 
    // same with the token to serialize printing
    c.printAccess <- true
 
    // run goroutine to crawl top level url.
    // since we are starting exactly one goroutine here, we wait
    // for a single completion report.  receipt means that all
    // lower levels have also completed and it is safe to return
    // --and allow the caller to return, in this case, main().
    done := make(chan bool)
    go c.c2(url, depth, done)
    <-done
}
 
func (c *crawlShared) c2(url string, depth int, pageDone chan bool) {
    // the function has multiple return points.  all of them must
    // report goroutine completion by sending a value on pageDone.
    if depth <= 0 {
        pageDone <- true
        return
    }
 
    body, urls, err := c.Fetch(url)
    if err != nil {
        // here's how to print:
        // take the token (waiting for it if it's not there.)
        <-c.printAccess
        // do whatever you need to do while other goroutines are
        // excluded from printing.
        fmt.Println(err)
        // put the token back, allowing others to print again.
        c.printAccess <- true
 
        pageDone <- true
        return
    }
 
    // same sequence of steps to print found message
    <-c.printAccess
    fmt.Printf("found: %s %q\n", url, body)
    c.printAccess <- true
 
    // "found" means the url was fetched without error and that urls
    // on the fetched page are collected in the slice "urls."
    // synchronization to crawl these urls in parallel is implemeted
    // with the uDone channel.  create the channel, count the number of
    // goroutines started, then wait for exactly that many completions.
    uDone := make(chan bool)
    uCount := 0
 
    // salt shaker pattern for map access:  get the map from
    // the channel, and then hold it while iterating over urls.
    // this works with the assumption that all of the operations here
    // take trivial time compared to the relatively lengthy time
    // to fetch a url.  other than map access (which is what we need
    // exclusive access for!) the only operations are iterating over
    // a string slice, incrementing an integer, and starting
    // a goroutine.  these all run very fast so it is reasonable and
    // best to hold "the lock" that is, exclusive map access, while
    // running through this loop.
    m := <-c.mapAccess
    for _, u := range urls {
        if !m[u] {
            m[u] = true
            uCount++
            go c.c2(u, depth-1, uDone)
        }
    }
    c.mapAccess <- m
 
    // wait for the number of goroutines started just above.
    for ; uCount > 0; uCount-- {
        <-uDone
    }
 
    // and finally, report completion of this level
    pageDone <- true
}

func main() {
    Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
    if res, ok := (*f)[url]; ok {
        return res.body, res.urls, nil
    }
    return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
    "http://golang.org/": &fakeResult{
        "The Go Programming Language",
        []string{
            "http://golang.org/pkg/",
            "http://golang.org/cmd/",
	    "http://golang.org/burek/",
        },
    },
    "http://golang.org/pkg/": &fakeResult{
        "Packages",
        []string{
            "http://golang.org/",
            "http://golang.org/cmd/",
            "http://golang.org/pkg/fmt/",
            "http://golang.org/pkg/os/",
        },
    },
    "http://golang.org/pkg/fmt/": &fakeResult{
        "Package fmt",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
    "http://golang.org/pkg/os/": &fakeResult{
        "Package os",
        []string{
            "http://golang.org/",
            "http://golang.org/pkg/",
        },
    },
}

