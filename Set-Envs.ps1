$Env:TWILIO_ACCOUNT_SID = (Get-Secret -Name TWILIO_ACCOUNT_SID -AsPlainText)
$Env:TWILIO_AUTH_TOKEN = (Get-Secret -Name TWILIO_AUTH_TOKEN -AsPlainText)

pwsh.exe -noexit -command 'try { . "c:\Users\jackh\AppData\Local\Programs\Microsoft VS Code\4fe60c8b1c\resources\app\out\vs\workbench\contrib\terminal\common\scripts\shellIntegration.ps1" } catch {}'