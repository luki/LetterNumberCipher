#Numcipher

This repository includes simple letter-number enciphering & deciphering algorithms written in Swift 3.0 and Go. Everyone is
**allowed to implement** them into any project, may it be **comercially or not**.

##Background

Back at the start of the year, I have sent a cryptic message to my girlfriend that used the (easy) letter-number code mentioned before. She got it right away, but at first, I told her her idea of it being the letter-number code was wrong (All though she was obviously right).

Today, we both do sometimes send these messages, since we both understand how to read them and have the tools to easily encipher and decipher messages.

##Execution

Two files, including algorithms for enciphering/encrypting and deciphering/decrypting for both Swift 2.0 and Go can be found in /src, in their respective folders, named after what they are meant to be used for. Added to a project, function(s) can be called and a message, wanted to be enciphered/encrpted entered.

###Swift

**Note:** The functions written in Swift require helper functions in order to work. These are extensions to the following typtes:

* String
  * toNumber: Returns an Integer that represents the letter's position in the alphabet.
  * loopValue: Returns an Integer that can be used for going through a string, for example in a for-loop (=string.characters.count - 1)
* Int
  * toLetter: Returns a string (=letter) that is found at the position in the alphabet.

####Examples

```swift
encipher("Hello")
```

```swift
decipher("8-5-12-12-15")
```

###Go

Go requires to have it [installed](https://golang.org/dl/) on the machine it should be run on. If it is, act as follows:

1. Select a folder for placing files `cd /Users/username/Documents`
2. Clone this repository: `git clone https://github.com/luki/numcipher.git`
3. Select the Go folder in local repository: `cd /Numcipher-master/src/go`
4. Run a program: `go run encryption.go`

####Examples

```go
encryption("Hello")
```

```go
decryption("8-5-12-12-15")
```


=================

##Feedback
For feedback, please send an email to lamuller@protonmail.com
