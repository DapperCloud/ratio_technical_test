# ratio_technical_test

## Foreword
I'm not a Go developer (yet) so this is probably not idiomatic Go (disclaimer: it might even look like Java a bit); 
hopefully, it doesn't look too horrible, still!

I'd like to add more integration tests (right now they're only for error cases), but it would take me quite some time: 
since the program output is non deterministic because of randomness and map sorting, I'd have to either do some smart parsing, or change the code to get deterministic output for testing.

I was told this should only take a few hours, so I decided not to overdo it too much. :) 

## Constraints and assertions
- A map cannot be empty (at least one city)
- All cities must have at least one road
- A city cannot have more than one road in a given direction
- A city cannot have more than one road to another city

There are no other constraints. Which means that you may well have:
```
A north=C
B north=C
C east=A
```
"A has a north road to C" does not mean "C has a south road to A", not even "C has a road to A". It's kind of a free graph.

You can see it as a set of cities, with roads that can be curved and crossing each other. "North road" doesn't mean that the road goes towards north, only that the road *starts* from the north side of the city.

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