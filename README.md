# vintexport

Golang scrapper to extract images from a vinted post.

> [!WARNING]
> This program may only be used for your personal posts. Scrapping other's people posts might get you banned.

## Features

- **Image validation** - Images are validated before being downloaded
- **Performance** - Lightning fast performance
- **Lightweight** - Minimal dependencies

## Compiling

> [!NOTE]
> Before continuing, make sure you have golang installed and fully working.

Edit the makefile to target your platform:

```Makefile
# Makefile
GOOS   ?= linux # linux / windows / darwin
```

Compile the program using make:

```bash
make build
```

## Usage

1. Run `vintexport` in your terminal
2. Input the post url

## Troubleshooting

### Common Issues

#### Photos location

**Problem:** Exported photos cannot be found.

**Solution:** Check the path where the program was ran, your photos should be there.

### Getting Help

If you're still having issues:

1. Search existing [issues](https://github.com/rotwo/vintexport/issues)
2. Create a new issue with detailed information
