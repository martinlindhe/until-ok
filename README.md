# until-ok <command>

Keeps retrying until given command works

```
~ ❯ until-ok ssh host
ssh: Could not resolve hostname host: nodename nor servname provided, or not known
ssh: Could not resolve hostname host: nodename nor servname provided, or not known
ssh: Could not resolve hostname host: nodename nor servname provided, or not known
...
...
Linux host 5.16.0-6-amd64 #1 SMP PREEMPT Debian 5.16.18-1 (2022-03-29) x86_64
~ >
```
