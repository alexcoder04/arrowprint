
# arrowprint

Use Arch Linux's makepkg-like output in your Go programs!

Windows support through [mattn/go-colorable](https://github.com/mattn/go-colorable).

<img style="width: 100%;" src="./screenshots/example01.png" alt="example screenshot">

### Code

```go
arrowprint.InfoC("Starting program...")
arrowprint.Suc0("Initializing system")
arrowprint.Suc1("checking variables")
arrowprint.Warn1("some variables are not defined, this may lead to errors")
arrowprint.InfoC("Running...")
arrowprint.Err0("An error occured, shutting down")
```

