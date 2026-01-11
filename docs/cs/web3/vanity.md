# Vanity

BTC靓号就是暴力生成，然后进行匹配特定的字符串

## 匹配类型

+ 前缀最快
+ 后缀次之
+ 正则最慢

## 开源库

+ [profanity](https://github.com/johguse/profanity)这个库有[漏洞](https://medium.com/amber-group/exploiting-the-profanity-flaw-e986576de7ab)
+ [vanitygen-plusplus](https://github.com/10gic/vanitygen-plusplus)
+ [VanitySearch](https://github.com/JeanLucPons/VanitySearch)
+ [Trongo](https://github.com/yl108305/Trongo)
+ [profanity2](https://github.com/1inch/profanity2)
+ [vanitygen](https://github.com/samr7/vanitygen)
+ [profanity-ether](https://github.com/ykky0/profanity-ether)

### vanitygen

这个库支持前缀和正则匹配

前缀GPU模式在4090/24GB上运行，差不多是600M H/s，GPU计算然后CPU正则匹配的话不到1M

```sh
./vanitygen++ -C LIST
./vanitygen++ 1Love
./vanitygen++ -F p2wpkh bc1qqqq
./vanitygen++ -F p2tr bc1pppp
./vanitygen++ -C ETH 0x999999
./vanitygen++ -C ETH -F contract 0x999999

./oclvanitygen++ -F compressed -k -i 19y73CRa -o 19y73CRa.txt
./vanitygen++ -F compressed -k -i 19y73CRa -o 19y73CRa.txt -t 44
./oclvanitygen++ -v -Z AAAA 1
./oclvanitygen++ -v -Z AAAA -l 14 1

./oclvanitygen -k -F contract -o contract0x900000000.txt -C ETH 0x900000000

./oclvanitygen -v -k -F contract -o contract0x900000000.txt -C ETH 0x900000000
./oclvanitygen -v -p 0 -k -F contract -o contract0x900000000.txt -C ETH 0x900000000

./vanitygen++ -F compressed -Z 0000000000000000000000000000000000000000000000000000000000000000 -l $((256-6)) 1PitScNLyp2HCygzad
./vanitygen++ -F compressed -Z 0000000000000000000000000000000000000000000000000000000000000000 -l $((256-20)) 1HsMJxNiV7TLxmoF6u


Usage: ./vanitygen++ [-vqnrik1NT] [-t <threads>] [-f <filename>|-] [<pattern>...]
Generates a bitcoin receiving address matching <pattern>, and outputs the
address and associated private key.  The private key may be stored in a safe
location or imported into a bitcoin client to spend any balance received on
the address.
By default, <pattern> is interpreted as an exact prefix.

Options:
-v            Verbose output
-q            Quiet output
-n            Simulate
-r            Use regular expression match instead of prefix
              (Feasibility of expression is not checked)
-i            Case-insensitive prefix search
-k            Keep pattern and continue search after finding a match
-1            Stop after first match
-a <amount>   Stop after generating <amount> addresses/keys
-C <altcoin>  Generate an address for specific altcoin, use "-C LIST" to view
              a list of all available altcoins, argument is case sensitive!
-X <version>  Generate address with the given version
-Y <version>  Specify private key version (-X provides public key)
-F <format>   Generate address with the given format (pubkey, compressed, script)
-P <pubkey>   Use split-key method with <pubkey> as base public key
-e            Encrypt private keys, prompt for password
-E <password> Encrypt private keys with <password> (UNSAFE)
-t <threads>  Set number of worker threads (Default: number of CPUs)
-f <file>     File containing list of patterns, one per line
              (Use "-" as the file name for stdin)
-o <file>     Write pattern matches to <file>
-s <file>     Seed random number generator from <file>
-Z <prefix>   Private key prefix in hex (1Address.io Dapp front-running protection)
-l <nbits>    Specify number of bits in prefix, only relevant when -Z is specified
-z            Format output of matches in CSV(disables verbose mode)
              Output as [COIN],[PREFIX],[ADDRESS],[PRIVKEY]
```
