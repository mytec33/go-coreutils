# coreutils - whoami

A re-implementation of the coreutils `whoami` program written in Go. This is inspired by the [Ubuntu uutils](https://github.com/uutils/coreutils) project written in Rust.

This tool hasn't been tested against the GNU coreutils test suite of the same name, but I intend to at some point.

## Example Usage

Return the effective user ID of the currently logged in user:
```
whoami
```

Return the effective user ID after the change in user ID by the `sudo` command:
```
sudo whoami
```

Note that on many platforms `whoami` is deprected in favor of the `id` command. The above would be run by issuing `id -un` in place of `whoami`.

## Platforms

| Platform            |Version| Builds|Tests|Architecture|
| :---------------------|-|:----:|:---:|:----------|
| `macOS (darwin)`      |15.x| ✅    |✅   |`arm64`      |
| `linux (Ubuntu)`|24.04| ✅    |✅   |`amd64`, `arm64`|
| `openbsd`             |7.6| ✅    |✅   |`armd64`     |

