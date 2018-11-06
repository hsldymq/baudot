# 博多码(Baudot Code)及其各种变体的编码解码工具

博多码(Baudot Code)是19世纪下半叶至20世纪上半叶被广泛用于电传打字机的一种字符集,由法国人Émile Baudot在1870年发明.

原版的博多码在早期在英国被推广使用,但是真正被大量普及是Donald Murray对电传打字机传输消息的改良由此在1901年对Baudot Code的改良,他的改良版本被成为Baudot-Murray Code,此后被标准化为International Telegraph Alphabet No.2(ITA2, 原版为ITA1), 它算得上是ASCII码的前身.

这个库是我无聊时候写的,实现了对ITA1和ITA2(standard及USTTY变体)编解码的功能.

#### Example

```golang
import (
    "github.com/hsldymq/baudot"
)

func main() {
    codec := baudot.newITA1(false)  // true:包含无效数据编解码数据将忽略, false:有无效数据会产生error
    // baudot.newITA2(false)
    // baudot.newUSTTY(false)

    codes, error := codec.Encode("X&Y")    // 编码消息为字节数组
    if error {
        // handle error
    }

    //////
    message, error := codec.Decode(codes)  // 解码博多码为消息字符串
    if error {
        // handle error
    }

}
```
