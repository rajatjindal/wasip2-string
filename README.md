- ensure you are using tinygo version with wasip2 support
- download [fermyon/spin](https://github.com/fermyon/spin)
- run `spin build && SPIN_VARIABLE_SECRET=works spin up`
- open new terminal and run: `curl http://127.0.0.1:3000`
- observe it returns following:

```
{
	"keys":["","",""],
	"variables":"\u0000\u0000\u0000\u0000\u0000"
}
```

The number of items in `keys` above is three, because we store three keys from within the code in `main.go`
The number of `\u0000` above is five, because the value of `SPIN_VARIABLE_SECRET` is `works`, which has 5 characters.

