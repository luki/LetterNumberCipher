#Anshouka
暗証化 (あんしょうか) - Anshōka: encryption; coding; password

Some time ago, I have sent a cryptic message to my girlfriend. She got the trick behind it really quickly, still do we sometimes use it out of fun.

Passed in letters will be translated to their position in the alphabet, meaning this tool does support just A-Z, but also German umlauts like 'Ä', 'Ö' and 'Ü', which will be translated to 'ae', 'oe' and 'ue'. Characters like 'ō' & ū from Romaji (Japanese with latin characters), are not supported yet.

Unkown characters will be translated to '0', since there is no letter '0'. In the translation back to original, 0 will be translated to '?'.  

What is included?
=================

Here you can find the encryption & decryption algorithms written in Go, but also the encryption algorithm written in Swift. The Swift variant uses own-made extensions, such as computed properties like toNumber, to turn letters (of type String) to numbers (of type Int), toLetter (the other way around), amount (instead of characters.count) and amountForIndex (amount of characters for loops and so on).

How can I run it?
=================

1. Install [Go](https://golang.org/dl/) on your computer
2. Clone this repository: `git clone https://github.com/miyohiki/alphanumeric.git`
3. Select the local repository: `cd alphanumeric`
4. Run a program: `go run encryption.go'`

=================

Ideas for the extension of the characters please to lukasamueller@icloud.com.