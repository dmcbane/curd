# CURD

***C**hange to one of a **U**ser's **R**ecurrent **D**irectories*

CURD allows you to have quick access to directories that you visit often without having to type the entire path.  If you've used tools like [autojump](https://github.com/wting/autojump), [fasd](https://github.com/clvv/fasd), or [z](https://github.com/rupa/z), then you get the idea.  It is the latest evolution of a script that I have used for years to jump from project directory to project directory.  I decided to rewrite it in [Go](https://golang.org/) as a learning exercise and to simplify the script so that it is easy to maintain across multiple operating systems.  The result is a tool that is easy to integrate into the Windows command prompt, Windows Powershell, Mac Terminal, or Linux/Unix sh/bash/zsh shells.

The name CURD?  It is one character off from curr which is the name of the original script that I used to change the current directory to one that I had previously saved.  The acronym is completely contrived to match.

## Installation

CURD is written in [Go](https://golang.org/) so you'll need to have it installed to build it.  Once [Go is installed](https://golang.org/doc/install) and GOROOT is added to your path, the following command will install CURD.

    go get -u github.com/dmcbane/curd

## Integration

To actually make CURD useful, it needs to be integrated into the terminal/command shell of your choice:

**Windows Command Prompt:** Copy [curr.bat](https://raw.githubusercontent.com/dmcbane/curd/master/curr.bat) into a folder in your path.

    @echo off
    curd %* > %TEMP%\vv.tmp
    set /p VV=<%TEMP%\vv.tmp
    cd /D "%VV%"

**Windows Powershell:** Add the contents of [curr.ps1](https://raw.githubusercontent.com/dmcbane/curd/master/curr.ps1) into your $profile.  This will create a function that will change to the directory specified by CURD.

    Function Get-Curd-Directory {
      [CmdletBinding()]
        Param($arg)
          $content = if ($arg) {curd $arg} Else {curd}
          Set-Location "$content"
    };Set-Alias curr Get-Curd-Directory -Description "Change the current directory to the selected curd directory."

**Mac/Unix/Linux sh, bash or zsh:** Add the contents of [curr.sh](https://raw.githubusercontent.com/dmcbane/curd/master/curr.sh) into your .profile, .bashrc or .zshrc to create a function that will change to the directory specified by CURD.

    function curr() {
      D=$(curd "$@")
      cd "${D}"
    }


### Bash Completion

If you have bash completion installed on your system, you can source the src/github.com/dmcbane/curd/curd_completion.bash file from your ~/.bashrc file to enable bash completion for curr/curd.

ZSH users can use the bash completion file by adding the following to their ~/.zshrc file.

```
autoload bashcompinit
bashcompinit
source ~/go/src/github.com/dmcbane/curd/curd_completion.bash
```

## How It Works

Curd allows you to save the current working directory or a specified path by keyword or to the default keyword. You can later retrieve the path by the same keyword or using the default.  (The default is indicated by not specifying a keyword.) Since changing the directory from within an application doesn't persist once the application exits, Curd can't directly change the current working directory.  To get around this, Curd just prints the directory to STDOUT so that we can use it to change the directory i.e.

     cd $(curd)

will change to the default path or

     cd $(curd go)

will change to the path stored by keyword go.  That is way too much typing for normal use, so the integrations mentioned above reduce it to the following:

     curr
or

     curr go

## Usage

My typical workflow starts with setting up the saved paths in CURD:

    cd ~/src/DMS
    curd save dms
    cd ~/GoogleDrive/dev/rust/projects/testing
    curd save rust --dir=~/GoogleDrive/dev/rust/projects
    curd save test
    cd ~/go/src/github.com/dmcbane/curd
    curd save
    curd save curd

Once the paths are saved, it's easy to bounce around from project to project:

    curr dms
    <work in dms project folder>
    curr
    <work in curd project folder>
    curr test
    <work work in test project folder>

You can use the list command to see what paths are defined:

    curd list

or the clean command to remove non existant paths that are defined:

    curd clean

Typing `curd --help` will display the help screen for CURD which lists all available commands.

```
CURD - Change to a User's Recurring Directory 1.2.2
H. Dale McBane<h.dale.mcbane@gmail.com>
Save and return to paths you visit often.

Usage:
    curd clean [--config <file>] [--verbose]
    curd (completion | comp) CMDLINE ...
    curd (ls | list) [-k | --keywords-only] [--config <file>] [--verbose]
    curd (rm | remove) [KEYWORD] [--config <file>] [--verbose]
    curd save [KEYWORD] [--dir <directory>] [--config <file>] [--verbose]
    curd (help | -h | --help)
    curd (version | -V | --version)
    curd [KEYWORD] [--config <file>] [--verbose]

Options:
    --config=<file>  Specify configuration filename [default: /home/hdalemcbane/.curdrc].
    --dir=<directory>  Specify path name to associate with keyword [default: <current directory>].
    -k, --keywords-only  Don't include the path names in the list command.
    -h, --help     Show this screen.
    -V, --version  Show version.
    -v, --verbose  Display extra information.

Examples:
    List all of the keywords and paths defined in the default configuration file.
        curd ls

    List all of the keywords defined in the default configuration file.
        curd ls -k

    List all of the paths in a specified configuration file.
        curd list --config some_configuration_file

    Clean paths from the default configuration that do not exist in the
    filesystem.
        curd clean

    Read the default path from the default configuration file.
        curd

    Save the current directory as the default path in the default configuration
    file.
        curd save

    Save the specified directory as the specified path in the default
    configuration file.
        curd save curd --dir=~/go/src/github.com/dmcbane/curd

    Remove the specified path from the default configuration file.
        curd remove essay

    Used by shell completion scripts.
        curd comp curd ls -

```


