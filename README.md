Giggalo
=======

Command Line Parser For Go

Giggalo is a command line options parsing package for the Go programming language. Currently, Go's standard library contains a package called flag, which is pretty basic.
Giggalo aims to do for Go what Apache's CLI library did for Java...aswell as having a funny name!

## Installation
Giggalo depends on a couple of my other packages for testing purposes. You only need to install these if you wish to run the test suite

    go get github.com/daviejaneway/C4G/src
    go get github.com/daviejaneway/C4GUnit/src
    
The Giggalo package itself can be installed with

    go get github.com/daviejaneway/Giggalo/src
    
## Example
Suppose we want to access the Unix-like command line arguments passed into our Go program. Giggalo can parse them into a neat bundle for us, like so:
```go
    func main() {
      //First, define the options we are looking for
      opts := Giggalo.OptGroup{
      	Options:[]Giggalo.Option{
      	  Giggalo.Option{Id:"r", Longid:"recursive"},
      	  Giggalo.Option{Id:"n", Longid:"lines", Consume:true}}}
      	  
      //Now we parse the actual end-user supplied arguments (would usually be os.Args)
      opts.Parse([]string{"-r", "-n", "200"})
      
      //Finally, we can query the OptGroup for our options
      r, r_err := opts.Get("r")
      n, n_err := opts.Get("n")
      
      //r = true, r_err = nil if the end-user supplied an -r flag
      //n == 200, n_err = nil if the end user supplied an -n flag
      
      //NOTE: Get(flag string) returns an error if either flag is not
      //passed in on the command line or a consuming Option is passed in without a value
    }
```
    
## Running the test suite
Giggalo has a test suite under Giggalo/tests. To build & run it, simply do this:

    git clone https://github.com/daviejaneway/Giggalo
    cd Giggalo
    go run tests/*.c4g.go
