# CAPTCHA

| pattern | left operand |  operator  | right operand |
|:-------:|:------------:|:----------:|:-------------:|
| `1 - 2` |   `1 - 9`    |  `1 - 3`   |    `1 - 9`    |
                
Number can be `1 - 9`  
Numeric string can be `"One" - "Nine"`  
Operator can be `"+", "-", "*"`

### Example
| pattern | left operand | operator | right operand |  Output  |
|:-------:|:------------:|:--------:|:-------------:|:--------:|
|    1    |       1      |    1     |       1       |  1 + One |
|    2    |       1      |    1     |       1       |  One + 1 |
|    1    |       2      |    2     |       1       |  2 - One |
|    2    |       2      |    2     |       1       |  Two - 1 |
|    1    |       1      |    3     |       1       |  1 * One |
