package awskms

import (
	"unsafe"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"

	"fmt"
	"os"
)

// BytesToString test
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes test
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// ReadAwsCred read aws credential from ./aws/credentials
func ReadAwsCred(arn string, m string) string {

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create KMS service client
	svc := kms.New(sess)

	// Encrypt data key

	// Replace the fictitious key ARN with a valid key ID
	keyID := arn

	// Encrypt the data
	result, err := svc.Encrypt(&kms.EncryptInput{
		KeyId:     aws.String(keyID),
		Plaintext: []byte(m),
	})

	if err != nil {
		fmt.Println("Got error encrypting data: ", err)
		os.Exit(1)
	}

	s := BytesToString(result.CiphertextBlob)
	
	return s
}

// DecryptAwsCred decrypt aws credential
func DecryptAwsCred(arn string, m string) string {

	s := StringToBytes(m)


    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    // Create KMS service client
    svc := kms.New(sess)

    // Decrypt the data
    result, err := svc.Decrypt(&kms.DecryptInput{CiphertextBlob: s})

    if err != nil {
        fmt.Println("Got error decrypting data: ", err)
        os.Exit(1)
    }

    blobstring := string(result.Plaintext)

    return blobstring
}