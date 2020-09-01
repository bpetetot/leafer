# Leafer

> It's a personal project not supposed to be maintained.
> But feel free to fork it or open PR if you want to improve it.

A personal mangas and comics manager and reader.

- Add or remove different library folders
- Simple reader for `zip`, `rar`, `cbr`, `cbz` comic files
- Extract metadata from [`Anilist API`](https://anilist.co/) or `ComicInfo.xml`
- Mark as read by chapters

## Getting started

### Stack

| App          | Description                                                   |
| ------------ | ------------------------------------------------------------- |
| `leafer-app` | CLI and HTTP Server written with `golang`                     |
| `leafer-web` | Wep application written with `react` using `create-react-app` |

### Prerequisites

- `golang` installed
- `node@12` installed

### Usage

Start the HTTP server:

```shell
cd leafer-app
go get          # Install dependencies
go run serve    # Start HTTP server with the CLI
```

Start the web application:

```shell
cd leafer-web
yarn            # Install dependencies
yarn start      # Start the web application
```

## Build the app binary

To have a final binary in the `./build` folder, launch the `Makefile`

```shell
make
```

Then to launch it:

```shell
cd build
./leafer serve
```