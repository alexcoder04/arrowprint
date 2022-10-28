
# arrowprint

Use Arch-Linux-like output in your Go programs!

## Features

 - Different levels and colors of messages
 - `printf`-like formatting support
 - Windows support through [mattn/go-colorable](https://github.com/mattn/go-colorable).

## Example

<img style="width: 100%;" src="./screenshots/example01.png" alt="example screenshot">

```go
arrowprint.InfoC("Starting program...")
step1 := "system"
arrowprint.Suc0("Initializing %s", step1)
arrowprint.Suc1("checking %s", step1)
arrowprint.Warn1("%s are not defined, this may lead to errors", "some variables")
arrowprint.InfoC("Running...")
arrowprint.Err0("An error occured, shutting down")
```

## Available functions

```go
// colon print, bold
// ` :: message`
func InfoC(msg string, args ...any) { }
func SucC(msg string, args ...any) { }
func WarnC(msg string, args ...any) { }
func ErrC(msg string, args ...any) { }

// level 0 print
// ` => message`
func Info0(msg string, args ...any) { }
func Suc0(msg string, args ...any) { }
func Warn0(msg string, args ...any) { }
func Err0(msg string, args ...any) { }

// level 1 print
// `   -> message`
func Info1(msg string, args ...any) { }
func Suc1(msg string, args ...any) { }
func Warn1(msg string, args ...any) { }
func Err1(msg string, args ...any) { }
```
