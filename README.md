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
`ZeroNet/data/HUBID/data/users/USERID`

### Convert to git commit messages

I have seen a lot of people use GitHub Issues as a blog system.
But I have not seen any one use git commit messages as a blog system.

zerome2md can convert ZeroMe posts to git commit messages.
Note that this is a one time export and an empty git repository is preferred.

```sh
mkdir /path/to/example
cd /path/to/example
git init

cd ZeroNet/data/HUBID/data/users/USERID
zerome2md -p=https://example.com/path -r=/path/to/example -n=YOUR_NAME -e=YOUR_EMAIL
```

`https://example.com/path` is where you put all of your zerome `*.jpg` files.
`YOUR_NAME` and `YOUR_EMAIL` will be used in the git commit message.

## License

0BSD
