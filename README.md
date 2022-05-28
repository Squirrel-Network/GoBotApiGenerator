<p align="center">
    <a href="https://github.com/Squirrel-Network/GoBotApi">
        <img width="300px" src="logo1.svg" alt="gopher logo" />
    </a>
    <br><br>
    <b>AutoGenerated Telegram BotAPI Framework for GoLang</b>
    <br>
    <a href="https://github.com/Squirrel-Network/gobotapi/tree/master/examples">
        Examples
    </a>
    •
    <a href="https://pkg.go.dev/github.com/Squirrel-Network/gobotapi#section-documentation">
        Documentation
    </a>
    •
    <a href="https://github.com/Squirrel-Network/gobotapi">
        Sources
    </a>
</p>

# GoBotAPI
[![pkg.go.dev](https://img.shields.io/github/v/tag/Squirrel-Network/gobotapi?color=00b7e0&label=go.pkg.dev)](https://pkg.go.dev/github.com/Squirrel-Network/gobotapi)
[![GO Version](https://img.shields.io/github/go-mod/go-version/Squirrel-Network/gobotapi)](https://go.dev/)
[![GitHub](https://img.shields.io/github/license/Squirrel-Network/gobotapi)](https://github.com/Squirrel-Network/gobotapi/blob/master/LICENSE.md)
![OS](https://img.shields.io/badge/platform-Linux%20%7C%20Windows%20%7C%20macOS-lightgrey)
![Architectures](https://img.shields.io/badge/architectures-x86__64%20%7C%20arm64v8%20%7C%20win__amd64%20%7C%20darwin__x64-blue)

> An elegant and modern BotAPI Framework for GoLang

``` go
package main

import "github.com/Squirrel-Network/gobotapi"
import "github.com/Squirrel-Network/gobotapi/types"
import "github.com/Squirrel-Network/gobotapi/methods"

func main() {
    client := gobotapi.NewClient("YOUR_TOKEN")
    client.OnMessage(func(message types.Message) {
        client.Invoke(&methods.SendMessage{
            ChatID: message.Chat.ID,
            Text:   "Hello World!",
        })
    })
    client.Run()
}
```

**GoBotAPI** is a modern and elegant AutoGenerated [BotAPI](https://core.telegram.org/bots/api) Framework. This Framework provides a pure
Go implementation **without any external libs**

In addition to the official API, this Framework also provides some **high-level
functions** that make it easier to use the API.

> The Telegram API scheme depends on your build, but if you don't want to compile by your self you
can use the package compiled from [pkg.go.dev/github.com/Squirrel-Network/gobotapi](https://pkg.go.dev/github.com/Squirrel-Network/gobotapi).

## How to install?
Here's how to add the GoBotApi Framework to your project, the command are given below:
``` bash
go get -u github.com/Squirrel-Network/gobotapi
```

<hr>
<p align="center">
    <a href="https://github.com/Squirrel-Network/GoBotApiGenerator">
        <img width="300px" src="logo2.svg" alt="gopher logo" />
    </a>
    <br>
    <br>
    <a href="https://github.com/Squirrel-Network/GoBotApiCompiler/releases">
        Releases
    </a>
    •
    <a href="https://github.com/Squirrel-Network/GoBotApiGenerator">
        Sources
    </a>
</p>

# GoBotAPI Generator
[![GO Version](https://img.shields.io/github/go-mod/go-version/Squirrel-Network/GoBotApiCompiler)](https://go.dev/)
[![GitHub](https://img.shields.io/github/license/Squirrel-Network/GoBotApiCompiler)](https://github.com/Squirrel-Network/GoBotApiCompiler/blob/master/LICENSE.md)
![OS](https://img.shields.io/badge/platform-Linux%20%7C%20Windows%20%7C%20macOS-lightgrey)
![Architectures](https://img.shields.io/badge/architectures-x86__64%20%7C%20arm64v8%20%7C%20win__amd64%20%7C%20darwin__x64-blue)

**GoBotAPI Generator** is the generator for the GoBotAPI Framework. It generates a new BotAPI Framework from the given
[BotAPI](https://core.telegram.org/bots/api) schema.

# How to use?
Just run the binary file from [releases](https://github.com/Squirrel-Network/GoBotApiCompiler/releases)
and choose the binary file you want to use on your operating system.

## Credits
Big thanks to [@Laky-64] for making this project possible, special thanks to [@geiccobs] for his own package as
starting point for this project, also thanks to [@empijei] for help about the project design and to
[@LucaTheHacker] for optimizations.

Just because the thanks aren't enough...
Huge thanks to [@MikiMleam] for the cool logo and to [@BluLupo] for helping to fix some stuff.

[@Laky-64]: https://github.com/Laky-64
[@geiccobs]: https://github.com/geiccobs
[@empijei]: https://github.com/empijei
[@LucaTheHacker]: https://github.com/LucaTheHacker
[@BluLupo]: https://github.com/BluLupo
[@MikiMleam]: https://t.me/MikiMleam
