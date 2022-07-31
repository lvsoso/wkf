# wkf

## Usage
> This cli template shows the date and time in the terminal

wkf

## Description

```
This is a template CLI application, which can be used as a boilerplate for awesome CLI tools written in Go.
This template prints the date or time to the terminal.
```
## Examples

```bash
wkf date
wkf date --format 20060102
wkf time
wkf time --live
```

## Flags
|Flag|Usage|
|----|-----|
|`--config string`|config file (default is $HOME/.wkf.yaml)|
|`--debug`|enable debug messages|
|`--disable-update-checks`|disables update checks|
|`--raw`|print unstyled raw output (set it if output is written to a file)|

## Commands
|Command|Usage|
|-------|-----|
|`wkf completion`|Generate the autocompletion script for the specified shell|
|`wkf date`|Prints the current date.|
|`wkf help`|Help about any command|
|`wkf shell`|A Interactive input entry.|
|`wkf time`|Prints the current time|
# ... completion
`wkf completion`

## Usage
> Generate the autocompletion script for the specified shell

wkf completion

## Description

```
Generate the autocompletion script for wkf for the specified shell.
See each sub-command's help for details on how to use the generated script.

```

## Commands
|Command|Usage|
|-------|-----|
|`wkf completion bash`|Generate the autocompletion script for bash|
|`wkf completion fish`|Generate the autocompletion script for fish|
|`wkf completion powershell`|Generate the autocompletion script for powershell|
|`wkf completion zsh`|Generate the autocompletion script for zsh|
# ... completion bash
`wkf completion bash`

## Usage
> Generate the autocompletion script for bash

wkf completion bash

## Description

```
Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(wkf completion bash)

To load completions for every new session, execute once:

#### Linux:

	wkf completion bash > /etc/bash_completion.d/wkf

#### macOS:

	wkf completion bash > /usr/local/etc/bash_completion.d/wkf

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion fish
`wkf completion fish`

## Usage
> Generate the autocompletion script for fish

wkf completion fish

## Description

```
Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	wkf completion fish | source

To load completions for every new session, execute once:

	wkf completion fish > ~/.config/fish/completions/wkf.fish

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion powershell
`wkf completion powershell`

## Usage
> Generate the autocompletion script for powershell

wkf completion powershell

## Description

```
Generate the autocompletion script for powershell.

To load completions in your current shell session:

	wkf completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion zsh
`wkf completion zsh`

## Usage
> Generate the autocompletion script for zsh

wkf completion zsh

## Description

```
Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions for every new session, execute once:

#### Linux:

	wkf completion zsh > "${fpath[1]}/_wkf"

#### macOS:

	wkf completion zsh > /usr/local/share/zsh/site-functions/_wkf

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... date
`wkf date`

## Usage
> Prints the current date.

wkf date

## Flags
|Flag|Usage|
|----|-----|
|`-f, --format string`|specify a custom date format (default "02 Jan 06")|
# ... help
`wkf help`

## Usage
> Help about any command

wkf help [command]

## Description

```
Help provides help for any command in the application.
Simply type wkf help [path to command] for full details.
```
# ... shell
`wkf shell`

## Usage
> A Interactive input entry.

wkf shell

## Description

```
A Interactive input entry.
```
# ... time
`wkf time`

## Usage
> Prints the current time

wkf time

## Description

```
You can print a live clock with the '--live' flag!
```

## Flags
|Flag|Usage|
|----|-----|
|`-l, --live`|live output|


---
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 31 July 2022**
