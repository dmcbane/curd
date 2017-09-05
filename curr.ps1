Function Get-Curd-Directory {
  [CmdletBinding()]
    Param($arg)
      $content = if ($arg) {curd $arg} Else {curd}
      Set-Location "$content"
};Set-Alias curr Get-Curd-Directory -Description "Change the current directory to the selected curd directory."

