# CURD

Change to one of a User's Recurrent Directories

CURD allows you to quickly jump to a directory without having to type the entire
path.  It is the latest evolution of a script that I have used for years to jump
between a couple of projects.  This evolution stems from my desire to try out a
simple project in go and the need to simplify my script so that it is easy to
maintain across multiple operating systems.  The result is quick and easy to use
with integration wrappers that make it easy to use from the Windows command
prompt or Powershell, from the Mac Terminal, and from Unix/Linux bash or zsh
terminals.

Installation

Use Git to clone this repository to the appropriate location within your $GOPATH
($GOPATH/src/github.com/dmcbane/ or %GOPATH%\src\github.com\dmcbane\), change to
 the curd directory, then use go to install curd.

     git clone https://github.com/dmcbane/curd.git
     cd curd
     go install

Ensure that your $GOPATH/bin folder is in your path and then integrate curr into
your terminal as follows:

Windows Command Prompt: Copy the curr.bat file into your $GOPATH/bin folder or
anywhere else in your path.

Windows Powershell: Copy the contents of the curr.ps1 file into your $profile
which will create a function that will change to the directory specified by
curd.

     notepad curr.ps1 $profile

Unix/Linux bash or zsh: copy the contents of curr.sh into your .bashrc or .zshrc
to create a function that will change to the directory specified by curd.

If you have difficulty getting it installed correctly, try checking out "How to
Write Go Code" at https://golang.org/doc/code.html


How It Works

Curd allows you to save the current working directory path by keyword or to the
default. You can later retrieve the path by the same keyword or the default
path.  Since changing the directory from within an application doesn't persist
once the application exits, Curd can't directly change the current working
directory.  To get around this, Curd just prints the directory to STDOUT so that
we can use it to change the directory i.e.

     cd $(curd)

to change to the default path or

     cd $(curd go)

to change to the path stored by keyword go.  Since that is way too much typing
for constant use, the functions or Windows batch file reduce usage to the
following:

     curr

or

     curr go

Curd can be used as follows:

     curd [-c|l|r|s] [keyword] where omitting the c, l, r, and s flags retrieves
     the path specified by the keyword or the default if the keyword is omitted.

     -c pathname - Select a configuration file to use instead of the default
     (~.curdrc).
     -l - List all of the paths saved in the configuration file.
     -r [keyword] - Remove the path specified by the keyword or the default path
     from the configuration file.
     -s [keyword] - Set the the path specified by the keyword or the default
     path.

