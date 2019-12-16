# chatbot-notifier

chatbot-notifier is a tool for sending nofification to telegram chat groups. <br/>
chatbot-notifier aims to make sending notification to telegram chat group in a secure and easy way through monitoring script.

## Usage Example

notifier send -f credential.yml -m message.txt <br/>
notifier encrypt -f credential.yml

## credential.yml format (SAMPLE NOT REAL INFOR)

```
aws:
- arn: arn:aws:kms:ap-southeast-1:XXXXXXXXXX:key/XXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXX
telegram:
- token: 2312312312:DASDASGSDFDSFADSA
  chatid: -32213123123
```