# catd

Polls for changes in some `foo.d/` directory and concatenates all files into one `foo` file.

## Usage

```
$ catd -help
Usage of catd:
  -dir="": Directory to poll.
  -file="": File to concatenate contents of dir to.
  -once=false: Run once and exit.

# Example: combine all ~/.ssh/config.d/* files into ~/.ssh/config and exit.
$ catd -dir ~/.ssh/config.d -file ~/.ssh/config -once
```
