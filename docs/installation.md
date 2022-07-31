# Quick Start - Install wkf

> [!TIP]
> wkf is installable via [instl.sh](https://instl.sh).\
> You just have to run the following command and you're ready to go!

<!-- tabs:start -->

#### ** Windows **

### Windows Command

```powershell
iwr instl.sh/lvsoso/wkf/windows | iex
```

#### ** Linux **

### Linux Command

```bash
curl -sSL instl.sh/lvsoso/wkf/linux | bash
```

#### ** macOS **

### macOS Command

```bash
curl -sSL instl.sh/lvsoso/wkf/macos | bash
```

#### ** Compile from source **

### Compile from source with Golang

?> **NOTICE**
To compile wkf from source, you have to have [Go](https://golang.org/) installed.

Compiling wkf from source has the benefit that the build command is the same on every platform.\
It is not recommended to install Go only for the installation of wkf.

```command
go install github.com/lvsoso/wkf@latest
```

<!-- tabs:end -->
