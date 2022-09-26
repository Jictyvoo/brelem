# Brazil Elements

Most of the brazilian elements, validators, formatters that are utilized by services written in Go.

## Division

This package has different packages for each functionality

- validators - Has all validators with the name of the document which will be validated
- brtime - Should work as `time` package, making possible to format, convert `time.Time`, `time.Weekday`, `time.Month`
  to Brazilian defaults.
- entities - Has an entity that can be used to store data, like a data-object that store a document info
