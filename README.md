# About

Searches a file for an agenda and prints it.

## Use

Example Bash profile use:

```bash
alias today="agenda -path $HOME/Journal/{YYYY}/{MM}-{DD}.md"
```

Example PowerShell profile use:

```powershell
function Get-JournalAgenda() {
    agenda.exe -path "$HOME/Journal/{YYYY}/{MM}-{DD}.md"
}
New-Alias today Get-JournalAgenda
```
