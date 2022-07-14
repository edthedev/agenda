# About

Searches a file for an agenda and prints it.

## Use

Example PowerShell profile use:

```powershell
function Get-JournalAgenda() {
    agenda.exe -path "$HOME/Journal/$(Get-Date -Format "yyyy/MM-dd").md"
}
New-Alias agenda Get-JournalAgenda
```
