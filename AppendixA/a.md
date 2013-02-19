附录A. 增强巴科斯-诺尔范式（ABNF）语法
==============================
本节提供了本文档中定义的元素按[RFC5234][RFC5234]记法的增强巴克斯诺尔范式（ABNF）的语法描述。下列ABNF用Unicode代码要点[W3C.REC-XML-20081126]的术语定义；这些字符通常以UTF-8编码。元素按首次定义的顺序排列。

一些定义遵循使用来自[RFC3986][RFC3986]“URI引用”的定义。

一些定义遵循使用这些通用的定义：

     VSCHAR     = %x20-7E
     NQCHAR     = %x21 / %x23-5B / %x5D-7E
     NQSCHAR    = %x20-21 / %x23-5B / %x5D-7E
     UNICODECHARNOCRLF = %x09 /%x20-7E / %x80-D7FF / %xE000-FFFD / %x10000-10FFFF
（UNICODECHARNOCRLF定义基于[W3C.REC-XML-20081126]2.2节中定义的字符，但忽略了回车和换行字符。）

- A.1. [“client_id”语法](a.1.md)
- A.2. [“client_secret”语法](a.2.md)
- A.3. [“response_type”语法](a.3.md)
- A.4. [“scope”语法](a.4.md)
- A.5. [“state”语法](a.5.md)
- A.6. [“redirect_uri”语法](a.6.md)
- A.7. [“error”语法](a.7.md)
- A.8. [“error_description”语法](a.8.md)
- A.9. [“error_uri”语法](a.9.md)
- A.10. [“grant_type”语法](a.10.md)
- A.11. [“code”语法](a.11.md)
- A.12. [“access_token”语法](a.12.md)
- A.13. [“token_type”语法](a.13.md)
- A.14. [“expires_in”语法](a.14.md)
- A.15. [“username”语法](a.15.md)
- A.16. [“password”语法](a.16.md)
- A.17. [“refresh_token”语法](a.17.md)
- A.18. [端点参数语法](a.18.md)

[RFC5234]:http://tools.ietf.org/html/rfc5234 "Augmented BNF for Syntax Specifications: ABNF"
[RFC3986]:http://tools.ietf.org/html/rfc3986 "Uniform Resource Identifier (URI): Generic Syntax"