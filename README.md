[![](https://images.microbadger.com/badges/image/nbarnum/memory-leak.svg)](https://microbadger.com/images/nbarnum/memory-leak "Get your own image badge on microbadger.com") [![](https://images.microbadger.com/badges/version/nbarnum/memory-leak.svg)](https://microbadger.com/images/nbarnum/memory-leak "Get your own version badge on microbadger.com") [![](https://images.microbadger.com/badges/commit/nbarnum/memory-leak.svg)](https://microbadger.com/images/nbarnum/memory-leak "Get your own commit badge on microbadger.com")

# docker-memory-leak

Go code pulled from [this gist](https://gist.github.com/bcotton/5d4c035846697e6a20a43d37cbb15937).

Docker container that leaks memory. Can be used to tune memory monitoring collection, thresholds, etc.

## Using

Make sure to set a memory limit `-m/--memory`:

```
$ docker run -m 134217728 --rm -it -p 8081:8081 nbarnum/memory-leak
```
