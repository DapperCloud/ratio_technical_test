# ratio_technical_test

This is my Go solution to Ratio's technical test!

I'm not a Go developer (yet) so this is probably not idiomatic Go (disclaimer: it might even look like Java a bit); 
hopefully, it doesn't look too horrible, still!

## Building and running

### Build
The solution doesn't have any external dependencies. To build it, all you have to do is clone it in your Go workspace and run:
```
$> make
```

It will build the solution into a binary called `monsters` in the `bin` subdirectory.

### Run

The command should be used as follows:

```
$> ./bin/monsters [path to world map file] [number of monsters to spawn]
```

## Testing

To run the unit tests, simply do:
```
$> make test
```

There are also a few integration tests. Run them with:
```
$> make integration-test
```