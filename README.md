# pocket-list-dedupe <!-- omit in toc -->

[![Build Status](https://github.com/kevinpollet/pocket-list-dedupe/workflows/build/badge.svg)](https://github.com/kevinpollet/pocket-list-dedupe/actions)
[![License](https://img.shields.io/github/license/kevinpollet/pocket-list-dedupe)](./LICENSE.md)

Remove duplicate items in your Pocket list.

## Table of Contents <!-- omit in toc -->

- [Install](#install)
- [Usage](#usage)
  - [Get a Consumer Key](#get-a-consumer-key)
  - [Commands](#commands)
- [Contributing](#contributing)
- [License](#license)

## Install

```shell
$ go get github.com/kevinpollet/pocket-list-dedupe
```

## Usage

### Get a Consumer Key

1. Go to https://getpocket.com/developer/apps/new
2. Enter App `name` and `description`
3. Select `Retrieve` and `Modify` permissions
4. Select `Desktop` platform
5. Accept the Terms of Service
6. Create Application
7. Copy and keep your `Consumer Key` secret

### Commands

```shell
Remove duplicate items in your Pocket list

Usage:
  pocket-list-dedupe [flags]

Flags:
  -c, --consumerKey string   Pocket application's Consumer Key
  -h, --help                 help for pocket-list-dedupe
```

## Contributing

Contributions are welcome!

Want to file a bug, request a feature or contribute some code?

1. Check out the [Code of Conduct](./CODE_OF_CONDUCT.md).
2. Check for an existing [issue](https://github.com/kevinpollet/pocket-list-dedupe/issues) corresponding to your bug or feature request.
3. Open an issue to describe your bug or feature request.

## License

[MIT](./LICENSE.md)
