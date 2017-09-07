# goh

> Quickly change your current working directory to a project in `$GOPATH/src/github.com/`

# Usage

Changes to the specified project directory

```
$ goh vutran/goh
```

Creates the directory (if necessary) and changes to the specified project directory

```
$ goh -m vutran/new-project
```

# Installation

1. Copy `.gohrc` and `source ~/.gohrc` in your shell startup script (`.bashrc`, `.zshrc`, etc.)
2. `go get -u github.com/vutran/goh`
3. `goh <username>/<repoName>`

# License

MIT © [Vu Tran](https://github.com/vutran/)
