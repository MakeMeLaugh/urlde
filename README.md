# URLde

Simple CLI tool to decode strings to or encode from urlencoded representation.

## Installation

This is a fun personal project.
There will be no binary pre-made binary releases or any other way to install `urlde` than:

```shell
go install gitlab.com/MakeMeLaugh/urlde@latest
``` 

## Usage

* Encoding:

```shell
~ urlde encode '<>'
input:   | '<>'
encoded: | '%3C%3E'
```

* Decoding

```shell
urlde decode '%3C%3E'
input:   | '%3C%3E'
decoded: | '<>'
```

## Contributing

There is no further development of this tool planned. It was made as a fun/helpful project.
If you'd like to expand it functionality - you are free to fork this repo and continue to work on your own copy.
