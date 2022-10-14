# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [2.0.0] - 2022-10-14
- Revert v1.1.0 change to `OrderLineItem` `ID` from type `json.Number` to `string`

## [1.1.1] - 2021-04-12
- Update test dependencies

## [1.1.0] - 2020-03-26
- Include custom user agent for debugging and informational purposes
- Fix issue with query params not being passed in GET/DELETE requests
- Change `OrderLineItem` `ID` to type `json.Number`

## [1.0.0] - 2019-09-13
- Initial release

[Unreleased]: https://github.com/taxjar/taxjar-go/compare/v2.0.0...HEAD
[2.0.0]: https://github.com/taxjar/taxjar-go/compare/v1.1.1...v2.0.0
[1.1.1]: https://github.com/taxjar/taxjar-go/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/taxjar/taxjar-go/compare/v1.0.0...v1.1.0
