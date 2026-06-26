$logo = "Dev Container - [$(($PSVersionTable).OS)]"
Write-Host "`e[2;32m$Logo`e[0m"

function Invoke-OhMyPosh {
  [Diagnostics.CodeAnalysis.SuppressMessageAttribute('PSAvoidUsingInvokeExpression', '', Justification = 'It is allowed in this scenario')]
  param()
  oh-my-posh init pwsh --config ~/.poshthemes/blue-owl.omp.json | Invoke-Expression
}

Invoke-OhMyPosh

Set-PSReadLineOption -PredictionViewStyle ListView

function HistoryFile {
  code (Get-PSReadLineOption).HistorySavePath
}
