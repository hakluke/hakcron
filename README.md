# hakcron
Easily schedule commands to run multiple times at set intervals (like a cronjob, but for a single command)

# Description
hakcron allows you to run a command at specific intervals. It was written with the intention of being able to quickly set up a cronjob in a tmux session or similar, without having to actually edit the crontab.

# Flags

`-c` is the command that you wish to run at set intervals.
`-f` defines the frequency of the command being run.


# Example

To run a command daily, you could do:

```
hakcron -f "daily" -c "curl hakluke.com/dostuff.php"
```

Similarly, to run hourly, you could do:

```
hakcron -f "hourly" -c "curl hakluke.com/dostuff.php"
```

`-f` can be set to yearly, montly, weekly, daily and hourly:

Entry                  | Description                                | Equivalent To
-----                  | -----------                                | -------------
yearly (or annually)   | Run once a year, midnight, Jan. 1st        | 0 0 0 1 1 *
monthly                | Run once a month, midnight, first of month | 0 0 0 1 * *
weekly                 | Run once a week, midnight between Sat/Sun  | 0 0 0 * * 0
daily (or @midnight)   | Run once a day, midnight                   | 0 0 0 * * *
hourly                 | Run once an hour, beginning of hour        | 0 0 * * * *

# Intervals
To be more specific, you can also use intervals like this:

- every 1h30m
- every 5s

For example:

```
hakcron -f "every 30s" -c "curl hakluke.com/dostuff.php"
```

# More Details
hakcron implements robfig's cron library, so for more details see here: 
https://pkg.go.dev/github.com/robfig/cron#hdr-Intervals

# Thanks
Huge thanks to [robfig](https://github.com/robfig) for writing the golang cron library.
