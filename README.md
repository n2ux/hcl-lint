HCL-LINT
---

Cloned from github.com/dwradcliffe/hcl-lint 

## Usage:

```sh
$ go get github.com/n2ux/hcllint
$ hcllint - < config_file
$ hcllint folder_path config_file ...
```

Lint's an arbitrary number of input arguments. If a folder path is specified,
it will check that folder for *.tf files and run the linter on them as well.
If the first argument is a '-' it will read from stdin and ignore any following arguments.