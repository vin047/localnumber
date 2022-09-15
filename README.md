# localnumber

Small command line application and Go library for formatting phone numbers in a local format.

Example:
- UK: +44 1932 567890 -> 01932567890
- India: +91 6297 062979 -> 6297062979

## Motivation
I needed a command line application that could take a phone number (formatted or unformatted) from any country and return it in a common format – the kind that you usually give to a local. This means without spaces and without the country code prefix, and also adhering to common practices.

E.g. In the UK nobody gives their number with the country code (+44) and we prefix our numbers with a '0'. In India, nobody gives their number with the country code (+91) but the '0' prefix is also commonly excluded.

This library and app uses https://github.com/nyaruka/phonenumbers, which is a Go port of https://github.com/google/libphonenumber.

## Features
- Removes all whitespaces.
- Removes the country code prefix.
- Removes the '0' prefix from Indian mobile phone numbers [^1] – it's not necessary [^2], and some places won't accept more than 10 digits [^3].

## Usage
### Binary
Download binary and run it, supplying first argument with a phone number. Use quotes to escape spaces.
```shell
$ ./localnumber "+91 6297 062979"
6297062979
```


### Library
```go
import "github.com/vin047/localnumber"

func foo() {
    input := "+91 6297 062979"
    number, err := localnumber.Format(input)
	if err != nil {
		panic(err)
	}
    print(number)   // 6297062979
}
```

## License
See [LICENSE](LICENSE) file.

[^1]: https://stackoverflow.com/questions/70520184/convert-an-international-phone-number-into-a-national-number

[^2]: https://en.wikipedia.org/wiki/National_conventions_for_writing_telephone_numbers#India

[^3]: https://www.jio.com/selfcare/recharge/mobility/
