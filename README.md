# Flog
> A simple file logger written in [Golang](https://golang.org/).

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://github.com/hackmac89/flog/graphs/commit-activity)
[![Golint code style](https://img.shields.io/badge/code_style-Golint-CFB69A.svg)](https://github.com/golang/lint)
[![Go Report Card](https://goreportcard.com/badge/github.com/golang/crypto)](https://goreportcard.com/report/github.com/hackmac89/flog)
[![codecov](https://codecov.io/gh/hackmac89/flog/branch/master/graph/badge.svg)](https://codecov.io/gh/hackmac89/flog)
[![Github all releases](https://img.shields.io/github/downloads/hackmac89/go-share/total.svg)](https://github.com/hackmac89/flog/releases/)

## Usage

Import package to your project.

Instantiate a new logger with a given filepath, where the logs shall be written to.

```go
logger, logErr := flog.NewFileLogger("test.log")
```

Call the appropriate function when logging **INFO**, **DEBUG**, **WARNING** or **ERROR** messages

```go
// print INFO message
logger.PrintInfo("Printing \"%s\" message", "INFO")

// print DEBUG message
logger.PrintDebug("Printing \"%s\" message", "DEBUG")

// print WARNING message
logger.PrintWarning("Printing \"%s\" message", "WARNING")

// print ERROR message
logger.PrintError("Printing \"%s\" message", "ERRROR")
```

### Example output

![example](https://raw.githubusercontent.com/hackmac89/flog/master/example.png)

## Release History

* 1.0.1 (05/20/19)
    * Added License
    * some changes in README.md

* 1.0.0 (05/20/19)
    * Initial commit/release

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **hackmac89** - *Initial work* - [hackmac89](https://github.com/hackmac89)

See also the list of [contributors](https://github.com/hackmac89/go-share/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Meta

hackmac89 – [@hackmac89](https://twitter.com/hackmac89) – hackmac89@filmdatenbank-manager.de – [https://github.com/hackmac89/](https://github.com/hackmac89/)

## Contributing

1. Fork it (<https://github.com/hackmac89/flog/fork>)
2. Create your feature branch (`git checkout -b feature/featurename`)
3. Commit your changes (`git commit -am 'Added some feature'`)
4. Push to the branch (`git push origin feature/featurename`)
5. Create a new Pull Request
