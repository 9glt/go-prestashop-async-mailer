# Send mails from prestashop in non-blocking way

Non blocking way to send emails

```
MAIL_HOST=mail.host.tld MAIL_USERNAME=user@host.tld MAIL_PASSWORD=password MAIL_FROM=user@host.tld MAIL_BINDTO=127.0.0.1:8888 ./app 
``` 

testing: copy send.php and edit "to" and "subject" also "content" parts.
```
php send.php
```
