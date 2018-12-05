# zerome2md - dump ZeroMe posts to markdown

This tool dumps ZeroMe posts to a markdown file,
so you can use it somewhere else
(for example, to be consumed by a static site generator,
or just for backup).

Currently only support dump posts,
not comments.

## Install

Clone the repository and compile it from source.
Then put the executable `zerome2md` to `PATH`.

## Usage

Go to your ZeroMe user directory,
then run `zerome2md > output.md`.

To render image files correctly,
you need to copy `*.jpg` files and `output.md` to the same destination directory.

### Where is your ZeroMe user directory?

Check your ZeroMe url, e.g.

http://127.0.0.1:43110/Me.ZeroNetwork.bit/?Profile/HUBID/USERID/yourid@zeroid.bit

Then you should run `zerome2md` under
`ZeroNet/data/HUBID/data/user/USERID`

## License

0BSD
