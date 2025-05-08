# Multilingual Greeting CLI – Go

What do we do when we start learning a new programming language? Yep, we all print "Hello " on our console. But this time, I'll take it further by not just printing "Hello World", I'll print "Hello Earth" using different earth spoken languages! 

## Features

- Written in Go using the `flag` package
- Supports multiple languages with a type-safe map
- Includes unit tests to make sure we're not messing things up

## Supported Languages

| Language | Code | Greeting             |
|----------|------|----------------------|
| English  | en   | Hello Earth          |
| French   | fr   | Bonjour le monde     |
| Burmese  | mm   | မင်္ဂလာပါ ကမ္ဘာမြေ       |
| Hebrew   | he   | שלום כדור הארץ         |
| Greek    | el   | Γεια σου Γη          |
| Latin    | lt   | Salve Terra          |

## Project Files

- `main.go` – Main application logic
- `main_test.go` – Unit tests to make sure we're saying "Hello Earth" correctly

## Behavior

If you give it a language code it doesnt know, it will politely tell you that the language is unsupported. No hard feelings, right?
