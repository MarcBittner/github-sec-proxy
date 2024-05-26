# github-sec-proxy

[![Software License][ico-license]](LICENSE.md)

This is project is a webproxy that allows for the creation of unique URLs that recreate google drive's "anyone with the link cans share" functionality for github private repos. This is a proof of concept and may have security implications I've not thought through. Use at your own risk

Please contact m@sec.technology if you have any questions

## Limitations

- I have not tested building and executing from a binary but it should "Just Work" (™)

## Structure

```
/main.go    - program logic
/env        - template for environment variables
/run.sh     - bash wrapper to parse envars from a file
/go.sum     - checksums for golang dependencies
/go.mod     - golang module properties definition
/README.md  - this file
```

## Setup

The following environment variables need to be set, as explained below:

```bash
GITHUB_ACCESS_TOKEN     -- Access token for the github account hosting the private repo
```

Copy the env template to .env and add the required values then source them into the proxy using the run.sh wrapper with the syntax "run.sh ${filename} ${ProgramExecution}", specifically "./run.sh .env go run main.go"

```bash

$ phaedrus@q.local: ~/gits/cw-messaging-api
[  git: main ] [ Exit: 0 ] [ last: 25.4ms ]$ ./run.sh .env go run main.go

```

## Usage

## TODO

## Security

If you discover any security related issues, please email m@sec.technology instead of using the issue tracker.

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.

[ico-license]: https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square
