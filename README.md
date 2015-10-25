# Submarine Cloud

A mobile game that is made with Unity3D and RoR, WebSocket server written in Go.

**NOTE: This repository does not include Asset Store assets.**

## Getting Started

Install tools.

```bash
$ brew tap shiwano/formulas
$ brew install robo musta
```

Make out `build.yml` from `build.example.yml`.

```bash
$ cp build.example.yml build.yml
$ vi build.yml # Edit variables.
```

Output tasks.

```bash
$ robo
```
