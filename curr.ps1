Function Get-Curd-Directory {
  [CmdletBinding()]
    Param($arg)
      $content = if ($arg) {curd $arg} Else {curd}
      Set-Location "$content"
};Set-Alias curr Get-Curd-Directory -Description "Change the current directory to the selected curd directory."

# Tab completion for curr: suggest saved keywords from `curd ls -k`.
Register-ArgumentCompleter -CommandName Get-Curd-Directory -ParameterName arg -ScriptBlock {
    param($commandName, $parameterName, $wordToComplete, $commandAst, $fakeBoundParameters)
    (curd ls -k) -split '\s+' |
        Where-Object { $_ -and $_ -like "$wordToComplete*" } |
        ForEach-Object {
            [System.Management.Automation.CompletionResult]::new($_, $_, 'ParameterValue', $_)
        }
}
