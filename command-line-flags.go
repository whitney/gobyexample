package main

import (
  "flag"
  "fmt"
)

// Command-line flags are a common way to specify options 
// for command-line programs. For example, in wc -l the -l 
// is a command-line flag.

// Go provides a flag package supporting basic command-line 
// flag parsing. We’ll use this package to implement our example 
// command-line program.
func main() {
  // Basic flag declarations are available for string, integer, 
  // and boolean options. Here we declare a string flag word with 
  // a default value "foo" and a short description. This flag.String 
  // function returns a string pointer (not a string value); we’ll 
  // see how to use this pointer below.
  wordPtr := flag.String("word", "foo", "a string")
  numbPtr := flag.Int("numb", 42, "an int")
  boolPtr := flag.Bool("fork", false, "a bool")

  // It’s also possible to declare an option that uses an existing 
  // var declared elsewhere in the program. Note that we need to pass 
  // in a pointer to the flag declaration function.
  var svar string
  flag.StringVar(&svar, "svar", "bar", "a string var")

  // Once all flags are declared, call flag.Parse() to execute the 
  // command-line parsing.
  flag.Parse()

  // Here we’ll just dump out the parsed options and any trailing 
  // positional arguments. Note that we need to dereference the 
  // points with e.g. *wordPtr to get the actual option values.
  fmt.Println("word:", *wordPtr)
  fmt.Println("numb:", *numbPtr)
  fmt.Println("fork:", *boolPtr)
  fmt.Println("svar:", svar)
  fmt.Println("tail:", flag.Args())
}

// To experiment with the command-line flags program it’s best 
// to first compile it and then run the resulting binary directly.
// $ go build command-line-flags.go

// Try out the built program by first giving it values for all flags.
// $ ./command-line-flags -word=opt -numb=7 -fork -svar=flag

// Note that if you omit flags they automatically take their default values.
// ./command-line-flags -word=opt

// Trailing positional arguments can be provided after any flags.
// $ ./command-line-flags -word=opt a1 a2 a3

// Note that the flag package requires all flags to appear before 
// positional arguments (otherwise the flags will be interpreted as 
// positional arguments).
// $ ./command-line-flags -word=opt a1 a2 a3 -numb=7

// Use -h or --help flags to get automatically generated help text 
// for the command-line program.
// $ ./command-line-flags -h

// If you provide a flag that wasn’t specified to the flag package, 
// the program will print an error message an show the help text again.
// $ ./command-line-flags -wat
