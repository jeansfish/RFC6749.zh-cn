附录B. 使用application/x-www-form-urlencoded媒体类型
==================================
在本规范公布的时候，“application/x-www-form-urlencoded”媒体类型在[W3C.REC-html401-19991224]的17.13.4节中定义但未在IANA MIME媒体类型注册表（<[http://www.iana.org/assignments/media-types](http://www.iana.org/assignments/media-types)>）中注册。此外，该定义是不完整的，因为它未考虑非US-ASCII的字符。
在使用这种媒体类型生成有效载荷使时为解决这个缺点，名称和值必须首先使用UTF-8字符编码方案[RFC3629]编码；作为结果的八位序列然后需要使用在[W3C.REC-html401-19991224]中定义的转义规则进一步编码。
当从使用这种媒体类型的有效载荷中解析数据时，由逆向名称/值编码得到的名称和值因而需要被视作八位序列，使用UTF-8字符编码方案解码。
例如，包含六个Unicode代码点的值

     (1) U+0020 (SPACE), (2) U+0025 (PERCENT SIGN),
     (3) U+0026 (AMPERSAND), (4) U+002B (PLUS SIGN),
     (5) U+00A3 (POUND SIGN), (6) U+20AC (EURO SIGN)
将被编码成如下的八位序列（使用十六进制表示）：

    20 25 26 2B C2 A3 E2 82 AC
然后在有效载荷中表示为：

    +%25%26%2B%C2%A3%E2%82%AC

[RFC3629]:http://tools.ietf.org/html/rfc3629 "UTF-8"