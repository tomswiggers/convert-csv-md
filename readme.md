# Readme

Convert a CSV file with a table to a MD file

## Example

```
 n        ,Description,Calculation
 -------- , ---------------------------------------- , -------------- 
1,Description 1,10
2,Description 2,25
3,Description 3,15
4,Description 4,10
5,Description 5,20
6,Description 6,30
 *TOTAL* ,                                          ,110
```

convert-csv-md --filename example/calculation.csv

This will convert the CSV to MD.

| n        | Description                              | Calculation    |
| -------- | ---------------------------------------- | -------------- |
| 1        | Description 1                            | 10             |
| 2        | Description 2                            | 25             |
| 3        | Description 3                            | 15             |
| 4        | Description 4                            | 10             |
| 5        | Description 5                            | 20             |
| 6        | Description 6                            | 30             |
| *TOTAL*  |                                          | 110            |

