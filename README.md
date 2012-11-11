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
http://img203.imageshack.us/img203/1835/omskbird1502.png
```

Installation:
-------------

a. If you are on Mac OS X, just grab this binary: [shack-bin-0.1.zip](https://github.com/downloads/ivanzoid/shack/shack-bin-0.1.zip)

b. If you use other system or wish to compile from source:

1. Install Go in your system:

   * Ubuntu/Debian: `sudo apt-get install golang-go`
   * Mac OS X HomeBrew: `brew install go`
   * [Other systems](http://golang.org/doc/install)

2. Build:
   `go build shack.go`

3. Place `shack` to a place you prefer, e.g. `~/bin`


Configuration
-------------

1. Get your imageshack.us [API key](http://stream.imageshack.us/api/)

2. Create file .shack.cfg in home directory with the following content:
```
{
	"key": "your imageshack.us API key"
}
```

If you wish that your uploaded files to be stored inside your imageshack.us account, you also may specify your username and password:
```
{
	"key": "your imageshack.us API key",
	"user": "your imageshack.us username",
	"password": "your imageshack.us password"
}
```

