# Preflight

Config management in go inspired by terraform (learning exercise).

Feel free to create Issues/PRs and learn with me!!

## What Preflight Is.

Preflight is a learning exercise to see what is involved in making a simple yet
capable Config Management System. It has a plugin api so that users can add 
custom functionality, it has a simple yet robust syntax built using HCL, and 
string interpolation, and it has a plan command to overview changes quickly and
easily before they happen.

## What Preflight Is Not.

Preflight is not a production system (yet) maybe it will be some day but for now 
the intentions are purely academic.

## Status

Preflight is still very early and most of the system is just ideas half layed 
out in code, I really would like to build this out in the open where everyone 
can see, learn, and collaborate. 

I urge anyone who is curious about how CMSs work or have a greate idea for a 
killer feature please contribute!

As for things todo here it is:

- [ ] Implement way to consistently Diff system and state files to generate plan
- [ ] Implement string interpolation and variable system
- [ ] Implement system for retrieving Data, Plugins from various storage mediums
- [ ] Implement system for remote/local execution
- [ ] Build reference/standard plugins

## Syntax (Subject to Change)

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

## Building

Currently plugins need to be built independently of the main binary this can be
done by running the following.

```
./build/build-plugin.sh
```
