# dirhost
Host a directory of your choice

## Quick start
```sh
# Default address: 127.0.0.1:8000
# Default path: .
# Default log: "" (empty means no logging to disk)

dirhost -a 0.0.0.0:8000 -p ~/Pictures -l ~/dirhost.log    # Serves your ~/Pictures folder over 0.0.0.0 on port 8000 and log to ~/dirhost.log
dirhost -a 0.0.0.0:8000 -p ~/Downloads                    # Serves your ~/Downloads folder over 0.0.0.0 on port 8000 without logging to disk
dirhort -a 0.0.0.0:8000                                   # Serves the folder you ran dirhost in on 0.0.0.0 on port 8000 without logging to disk
dirhost                                                   # Serves the folder you ran dirhost in on 127.0.0.1 on port 8000 without logging to disk
```

## Note
This tool serves the contents of the specified folder **INCLUDING** all the content in its subfolders!

## Options
```sh
Usage: dirhost [flags]

Flags:
  -h, --help                     Show context-sensitive help.
  -a, --addr="127.0.0.1:8000"    Enter address to serve.
  -p, --path="."                 Enter folder to serve.
  -l, --log=""                   Enter filename to write logs to. (leave empty for no log file)
```
