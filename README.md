# FizzBuzz REST server
The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by "fizz", all multiples of 5 by "buzz", and all multiples of 15 by "fizzbuzz". The output would look like this: `1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...`.

This implementation lets you customize the FizzBuzz rules by calling an endpoint.

You can refer to the OpenAPI documentation for more information (`/swagger/index.html`).

## Example
### Classic FizzBuzz
```shell
$ curl http://localhost:1323/fizzbuzz?limit=15&int1=5&int2=3&str1=Fizz&str2=Buzz
["1","2","Buzz","4","Fizz","Buzz","7","8","Buzz","Fizz","11","Buzz","13","14","FizzBuzz"]
```

### Fetching the most used configuration
```shell
$ curl http://localhost:1323/statistics
{"hits":3821,"parameters":{"limit":15,"int1":5,"int2":3,"str1":"Fizz","str2":"Buzz"}}
```
