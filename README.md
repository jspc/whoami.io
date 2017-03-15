whoami.io
==

my login name, even in sudo, as a service

Two components:

The server lives in the project root.

```bash
$ go build
$ ./whoami.io
```

This lives on port 8000


The client lives in `./client`:

```bash
 $ go build
 $ ./client
james.condron
```
