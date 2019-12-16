# chatbot-notifier

chatbot-notifier is a tool for sending message with telegram bot while securing chat id and token using AWS KMS.<br/>
chatbot-notifier aims to make sending notification to telegram chat group in an easy yet secure way.

## How it works

chatbot-notifier uses aws-sdk-go to implement the encryption and decryption operation. Thus you will require an amazon web services access key id and secret access key. Similar to using terraform or terragrunt, AWS_PROFILE have to be pass in the command line. Example will be "AWS_PROFILE=<PROFILE NAME> notifier send -f credential.yml -m textfile.txt". This profile read from <HOME>/.aws/credentials. If the access key id and secret access key are set as default profile in <HOME>/.aws/credentials, then AWS_PROFILE will not need to be pass. Example "notifier send -f credential.yml -m textfile.txt".

Example of <HOME>/.aws/credential

```aws
[default]
role_arn = arn:aws:iam::123456789012:role/testing
source_profile = default
role_session_name = OPTIONAL_SESSION_NAME

[profile project1]
role_arn = arn:aws:iam::123456789012:role/testing
source_profile = default
role_session_name = OPTIONAL_SESSION_NAME
```

For more information, see <https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/guide_credentials_profiles.html> 

notifier will then get key from AWS KMS to encrypt credential.yml. ENCRYPTED credential.yml WILL NOT BE ABLE TO DECRYPT THROUGH notifier. THIS IS TO ENSURE THAT TOKEN AND CHAT ID ARE SAFE IN THE SERVER. <br/>

notifier will then be able to send message using credential.yml. <br/>

## Usage Example

### Encrypt credential.yml
notifier encrypt -f credential.yml <br/>

### Message can only be send after encryption
notifier send -f credential.yml -m message.txt <br/>


## credential.yml format (SAMPLE NOT REAL INFOR) (File can be other name)

```go
aws:
- arn: arn:aws:kms:ap-southeast-1:XXXXXXXXXX:key/XXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXX
telegram:
- token: 2312312312:DASDASGSDFDSFADSA
  chatid: -32213123123
```

## message.txt (File can be other name)

Any free text file.
