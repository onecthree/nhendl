# nhendl: nhentai doujin downloader
A lightweight nhentai doujin downloader built using go. Download doujin direct to nhentai without cookie/auth account, simple usage, and clean output cli/terminal info. nhendl using goroutine to spawn each "thread"-like to download every page when using --async option, on default use; nhendl will using sync-mode.

## Usage:
nhendl require doujin id/code from nhentai.net website, you can clearly see the id/code [133769] from url like `https://nhentai.net/g/133769`
<br>
- For default usage:
  > `nhendl -d 133769`
- nhendl can print the full info of doujin, add the option:
  > `nhendl -d 133769 --fulldesc`
- When you want using async-mode to download the doujin, add the option:
  > `nhendl -d 133769 --async`
  - `--async` option require faster connection to significant result.
  - When nhendl downloading using async-mode, nhendl will limit goroutines spawn to 20 concurrent connection, is require to avoid Cloudflare limit connection from same IP/connection.
<br>

For more usage info, try `nhendl -h` on your terminal.

## Installation/Build:
Requirements:
- go 1.17.6 or greater.

Build:
1. Download repository using git:
    > `git clone https://github.com/onecthree/nhendl.git`
2. Open terminal, access the directory of nhendl and build binary/executable file using go:
    > `go build .`
3. When build success/finish, you can use direct to binary/executable file:
    > `./nhendl -d 133769`
    or
    > `nhendl.exe -d 133769`
