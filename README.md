## Functional HTML Library for Go

A library to help you write HTML layouts as regular Go functions, rather than using a template.

Heavily inspired by the way Elm does HTML output - in fact, bar a few tweaks to help suit Go's syntax a bit better, it's almost identical.

Currently being developed on an 'as needed' basis, so please feel free to add missing elements if you want/need them.

See the test file for a complete example of an HTML page written with `htmlfunc`

## Tests

Kind of tricky to test, but running `./test.sh` will generate an `index.html` and test it against the *vnu* html validator if you have it installed.

Inspecting the HTML is handy development crutch, and then visually checking over the file in a browser should help confirm the output is an intended.

- [VNU HTML Validator](https://github.com/validator/validator/)

