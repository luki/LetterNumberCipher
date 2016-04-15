#Jikaeru

Jikaeru contains a simple letter-number code written in Swift and Go. You can use them for fun things, as well as giving your friends (a little) harder time understanding you.  

Story
=================

Back at the start of the year, I have sent a cryptic message to my girlfriend that used the (easy) letter-number code mentioned before. She got it right away, but at first, I told her her idea of it being the letter-number code was wrong (All though she was obviously right).

Today, we both do sometimes send these messages, since we both understand how to read them and have the tools to easily encipher and decipher messages.

How it works
=================

You get two files for each Swift and Go, which contain functions for both enciphering and deciphering. Simply implement the files in your project, call the needed function and enter the fitting message. This can be applied to a lot of use-cases.
<br><br>
Swift examples:

```swift
encipher("Hello")
```

```swift
decipher("8-5-12-12-15")
```
<br>
Go examples:
```go
encryption("Hello")
```

```go
decryption("8-5-12-12-15")
```


How to run it
=================

In the Swift variant, you do not have to just grab the functions out of the project and implement them, instead you have to implement the whole files, since they contain own-made extensions:

String
<br><ul>
<li>toNumber: Returns an Integer that represents the letter's position in the alphabet.</li>
<li>loopValue: Returns an Integer that can be used for going through a string, for example in a for-loop (=string.characters.count - 1)</li></ul>

Int
<br><ul>
<li>toLetter: Returns a string (=letter) that is found at the position in the alphabet.</li></ul>


In the Go variant, you do of course need to [install Go](https://golang.org/dl/) on your machine first. After this follow these steps:

1. Select a folder you want the files to be in `cd /Users/username/Documents`
2. Clone this repository: `git clone https://github.com/miyohiki/jikaeru.git`
3. Select the go folder in local repo: `cd /Jikaeru-master/src/go`
4. Run a program: `go run encryption.go`


=================
Ideas for the extension of the characters please to lukasamueller@icloud.com.
