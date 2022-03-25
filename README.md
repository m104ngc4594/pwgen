# pwgen

## Overview
pwgen is a password generator for the command line written in Go. It works on all major desktop and server operating systems (Linux, MacOS, BSD, Windows). 

pwgen generates passwords based on a set of rules, which by default include all uppercase characters, lowercase characters, numbers, symbols, and each password has 20 length.

pwgen is derived from the gopass command which does not meet my needs for generated password strength, so for security reasons I developed my own pwgen. 

## How to Install

### Binary Install

If you want to use pwgen as your password generator, simply download from Releases and unpack the pwgen binaries. The pwgen binaries have no external dependencies.

Currently, we provide pre-built pwgen binaries for Windows, and Linux for x64 architectures.

### Build and Install the Binaries from Gitee or Github (Advanced Install)

#### Prerequisite Tools

- [Git](https://git-scm.com/)
- Go (pwgen 0.1.0 only builds with >= Go 1.18.)

#### Install from Gitee
```
go install gitee.com/m104/pwgen
```

#### Install from Github
```
go install github.com/m104ngc4594/pwgen
```

## Usage

### Generate password from the command line
```
pwgen
```
It will generate 60 passwords each line by default. 
