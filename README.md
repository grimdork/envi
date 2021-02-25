# envi
Enironment variable-aware INI reading.

## Example

```go
	ini, err := envi.Load("debug.ini")
	if err != nil {
		return nil, err
	}
	ini.ForceUpper()
	ini.GetEnvString("", "hostname")
```

This loads the file `debug.ini` and sets parsing to forcing uppercase envvars. If the variable `HOSTNAME` exists in the environment, it's returned, otherwise the file contents are returned. If all else fails, the variable type's zero value is returned (`""` for strings, `0` for integers, `0.0` for floats and `false` for booleans).

If `ForceUpper()` wasn't used, it would simply look for `hostname`instead.
