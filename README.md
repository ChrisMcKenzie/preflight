# preflight
Config management in go inspired by terraform (learning exercise).

**EXPERIMENT** 

This is very much a learning exercise for me and anyone interested in learning
how to build a config management tool in golang with a extendable plugin system

Feel free to create Issues/PRs and learn with me!!

```
# vim: set ft=hcl:

data "script" "dotfiles" {
  source = "http://github.com/chrismckenzie/dotfiles/install.sh"
}

task "homebrew" "install_vim" {
  name = "vim"
  state = "present"
}

task "file" "create_vimrc" {
  name = "${path.home}/.vimrc"
  source = "${path.module}/vimrc"
  state = "present"

  attrs {
    owner = "chrism"
    group = "staff"
    perrmissions = 777
  }
}

task "script" "run_dotfiles_install" {
  content = "${data.script.dotfiles}"
}
```
