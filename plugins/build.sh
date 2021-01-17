echo "编译插件..."
go build -buildmode=plugin -o ./mail/mail.so ./mail/mail.go
go build -buildmode=plugin -o ./webhook/webhook.so ./webhook/webhook.go
go build -buildmode=plugin -o ./sms/sms.so ./sms/sms.go