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

Make out `config.*.yml` from `config.example.yml`.

```bash
$ cp config.example.yml config.development.yml
$ vi config.development.yml # Write the build settings.
```

Output tasks.

```bash
$ rake -T
```
