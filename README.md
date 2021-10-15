# bech32

`bech32` implements a bech32 decoder suitable for decoding Cardano CIP-0019 addresses, 
https://github.com/cardano-foundation/CIPs/blob/master/CIP-0019/CIP-0019.md


```go
hrp, data, err := bech32.Decode("addr1vx2fxv2umyhttkxyxp8x0dlpdt3k6cwng5pxj3jhsydzers66hrl8")
...
```
