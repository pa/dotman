![Release with goreleaser](https://img.shields.io/github/workflow/status/pa/dotman/Release%20with%20goreleaser?label=Release%20with%20goreleaser&logo=GitHub&style=for-the-badge)
![GitHub Release](https://img.shields.io/github/v/release/pa/dotman?label=dotman%20release&logo=GitHub&style=for-the-badge)
![Go Version](https://img.shields.io/github/go-mod/go-version/pa/dotman?label=go%20version&logo=go&style=for-the-badge)

# dotman - dot(files) man(ager)

dotman is a [go](https://go.dev/) based simple and light weight tool for managing [dotfiles](https://en.wikipedia.org/wiki/Hidden_file_and_hidden_directory). This tool uses a [bare Git repository](https://www.atlassian.com/git/tutorials/dotfiles) that means your `$HOME` will be your git work tree. Also, it offers a plugin manager with rules to copy files from source git repo and even directory which most of the tools out there wouldn't support.

I had been inspired by [Bhupesh's](https://github.com/Bhupesh-V) project [dotman](https://github.com/Bhupesh-V/dotman) after which I have named this tool.

## Requirements
- [Git](https://git-scm.com/) executable insatlled on your machine
- A Git repo, where you version control your dotfiles

## Installation

Just a binary for now, so go to the [release page](https://github.com/pa/dotman/releases), find the version you want and download the archive. Unpack it and put the binary to somewhere you want (on UNIX-y systems, /usr/local/bin or the like). Make sure to turn on the executable bits if you are using custom location.

## Demo
A quick demo of the tool,

[![dotman demo video](assets/dotman-demo.gif)](assets/dotman-demo.gif "dotman demo video")

## Configuration File
Create a config file named `.dotman-config.<suported extention>` under your `$HOME` directory. The dotman supports various config file formats JSON, TOML and YAML.

.dotman-config.yaml
```
autoCommit: <boolean>
externals:
    <parentPath>:
        - url: <git url>
          paths:
            - <sourceDirFromGitRepo> <targetDir>
```

below is a example of dotman config

```
autoCommit: true
externals:
  .config/fish:
    - url: https://github.com/IlanCosman/tide
      paths:
        - completions completions
        - conf.d conf.d
        - functions functions
    - url: https://github.com/jethrokuan/z
      paths:
        - conf.d conf.d
        - functions functions
  .vim/pack/plugins/start:
    - url: https://github.com/preservim/nerdtree.git
```

## Usage

```
Usage:
  dotman [command]

Available Commands:
  add              Add file contents to the index
  commit           Record changes to the repository
  completion       Generate completion script
  config           Set and read git configuration variables
  diff             Show changes between commits, commit and working tree, etc
  help             Help about any command
  init             Clones dotfiles repo from remote git repository
  pull             Fetch from and merge with another repository or a local branch
  push             Update remote refs along with associated objects
  reset            Reset current HEAD to the specified state
  stash            Stash away changes
  status           Show the working tree status
  update-externals Downloads and updates git externals like plugins, etc
  ```

## Example
```
# dotman init
❯ dotman init
Using config file: /Users/pa/.dotman-config.yaml
git repo url <pass your dotfiles git repo url>

# dotman add
❯ dotman add ~/.zshrc
Using config file: /Users/pa/.dotman-config.yaml

# dotman status
❯ dotman status
Using config file: /Users/pa/.dotman-config.yaml
On branch main
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
	new file:   ../../../.zshrc

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
	modified:   ../../../.config/fish/config.fish

Untracked files not listed (use -u option to show untracked files)

# dotman push
❯ dotman push -u origin <branch name>

# dotman update-externals
❯ dotman update-externals
Using config file: /Users/pa/.dotman-config.yaml
Cloning into '/Users/pa/.dotman/externals/tide'...
remote: Enumerating objects: 6661, done.
remote: Counting objects: 100% (866/866), done.
remote: Compressing objects: 100% (173/173), done.
remote: Total 6661 (delta 778), reused 713 (delta 688), pack-reused 5795
Receiving objects: 100% (6661/6661), 6.16 MiB | 3.22 MiB/s, done.
Resolving deltas: 100% (4257/4257), done.
Cloning into '/Users/pa/.dotman/externals/z'...
remote: Enumerating objects: 582, done.
remote: Counting objects: 100% (112/112), done.
remote: Compressing objects: 100% (68/68), done.
remote: Total 582 (delta 57), reused 65 (delta 30), pack-reused 470
Receiving objects: 100% (582/582), 92.20 KiB | 1.74 MiB/s, done.
Resolving deltas: 100% (313/313), done.
```


## TODO
- Add more git native commands
- Write unit tests
- Add GitHub action to build and test
- Release package to various distributions