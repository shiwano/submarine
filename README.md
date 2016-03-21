# Submarine Cloud

A mobile game that is made with Unity3D and RoR, WebSocket server written in Go.

**NOTE: This repository does not include fee-charging assets of the Asset Store.**

## Getting Started

Install tools.

```bash
$ brew ruby node
$ gem install rake
$ npm install -g typhen
```

Make out `tools/build/config.*.yml` from `tools/build/config.example.yml`.

```bash
$ cd tools/build
$ cp config.example.yml config.development.yml
$ cp config.example.yml config.production.yml
$ vim -o config.development.yml config.production.yml # Write the build settings.
```

Output tasks.

```bash
$ rake -T
```
