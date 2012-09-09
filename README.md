What it is?
-----------
It is utility to upload your images to [imageshack.us](http://imageshack.us) with command line.

Example of use:
-----------
```
$ shack OmskBird150-2.png
```
Output:
```
http://img214.imageshack.us/img214/1835/omskbird1502.png
```

Installation:
-------------

1. Install Go in your system:

   * Ubuntu/Debian: `sudo apt-get install golang-go`
   * Mac OS X HomeBrew: `brew install go`
   * [Other systems](http://golang.org/doc/install)


2. Get your [API key](http://stream.imageshack.us/api/)

3. Place your API key to shack.go

4. Build:
   `go build shack.go`

5. Place `shack` to a place you prefer, e.g. `~/bin`
