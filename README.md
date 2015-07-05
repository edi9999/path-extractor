# PathExtractor (pe)

![demo](https://raw.github.com/edi9999/i/master/demo.gif)

PathExtractor is a command line tool that extracts a list of files/paths from stdin.

Advantages over [fpp](https://github.com/facebook/PathPicker):

 * It does only one thing : more unixy
 * You can use it with any fuzzy finder, such as [fzf](https://github.com/junegunn/fzf),[peco](https://github.com/peco/peco),[percol](https://github.com/mooz/percol),[pick](https://github.com/thoughtbot/pick),[selecta](https://github.com/garybernhardt/selecta/)
 * It doesn't wait for stdin to be finished to output the paths
 * It is faster
 * It is much smaller (easily understandable)
 * You can also use it without a fuzzy finder for programmatic usage

For example, you could write:

    git status | pe

to get a list of the files that were added/changed, without all the formating

One of the most common usage is to create an alias that will automatically run :

  `pe` + a command line fuzzy finder such as fzf + an action such as opening that file in your favorite editor.

For example, using `zsh` , I have as an alias:

    alias -g P='| pe | fzf | read filename; [ ! -z $filename ] && vim $filename'

With `bash`:

    bind '"PP": "| pe | uniq | fzf | while read filename; do [ ! -z $filename ] && </dev/tty vim $filename; done\n'

So that If I run

    `git status P`

or

    `git status PP`

to quickly open one of the changed files in vim

Other usage ideas:

With zsh:

    # Copy selected path to clipboard
    alias -g C='| pe | fzf | read filename; [ ! -z $filename ] && echo -n $filename | xclip -selection c'

With bash:

    bind '"CC": "| pe | fzf | read filename; [ ! -z $filename ] && echo -n $filename | xclip -selection c\n"'

# Installation

```
git clone # in your go path
go test
go build
go install
```
