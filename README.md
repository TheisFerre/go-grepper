# go-grepper

A small for-fun CLI tool written in GO to find lines in text using regex. The tool matches the behaviour of the linux command line tool `grep`.

## Usage

go-grepper can be used either from a linux pipe or from a file directly. It offers regex based pattern matching.

### Using go-grepper from a linux pipe

To use `go-grepper` to read from a linux pipe, the following command is available:

```bash
go-grepper match pipe <regex pattern>
```

As an example, of using `go-grepper` inside this git directory to find all files that start with `go`:

```bash
ls | go-grepper match pipe '^go'
```

The result can be seen in the picure below

![alt text](https://github.com/TheisFerre/go-grepper/blob/main/img/go-grepper-pipe.png)

### Using go-grepper to read a file

To use `go-grepper` to read from a file, the following command is available:

```bash
go-grepper match file <file name> <regex pattern>
```

As an example, we can use the file `lorem-imsum.txt` in this repository to find all the cases that matches the following regex pattern `[E|e]rat dolor`. Which means that the first letter E can be upper or lowercase and then it should match the rest exact.

```bash
go-grepper match file lorem-ipsum.txt '[E|e]rat dolor'
```

The result can be seen in the picure below

![alt text](https://github.com/TheisFerre/go-grepper/blob/main/img/go-grepper-file.png)