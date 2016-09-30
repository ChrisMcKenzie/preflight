# preflight
Config management in go inspired by terraform (learning exercise).

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
  path = "${path.home}/.vimrc"
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
