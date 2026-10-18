# granio-cross-images

Компонент независимого проекта **Granio — Гранио**.

**Исходный репозиторий:** [fyne-cross-images](https://github.com/fyne-io/fyne-cross-images).
**Цель форка:** Рецепты контейнеров кросс-компиляции Granio.

История изменений, лицензии и сведения об исходных авторах сохранены.
Go-модули переведены на собственные пути `github.com/Uzer007-admin/…`;
минимальная версия Go — **1.27.1**. Ядро и общие API подключаются из
[`granio/v2`](https://github.com/Uzer007-admin/granio).

Состояние переноса и результаты проверок приведены в
[реестре экосистемы](https://github.com/Uzer007-admin/granio/blob/dev/granio/docs/ECOSYSTEM.md).
Наличие форка само по себе не означает проверку всех платформ и сценариев.

Рецепт базового образа устанавливает опубликованные CLI Granio и собственный
помощник кросс-компиляции. Имя `fyne` внутри образа сохранено как ссылка на
`granio` для совместимости команд. Образы Granio ещё не собраны и не опубликованы;
при использовании upstream-образов их содержимое остаётся исходным Fyne.

---

## Документация исходного проекта

Команды установки и ссылки в сохранённом тексте ниже относятся к исходнику.
Для Granio используются новые пути модулей, перечисленные выше.

# Fyne Cross Images

fyne-cross is a simple tool to cross compile and create distribution packages
for [Fyne](https://fyne.io) applications using docker images that include Linux, 
Windows, FreeBSD, and a macOS SDK, along with the Fyne
requirements.

This project provides the docker images required by fyne-cross to cross compile.

## Requirements

- docker / podman

## Supported targets

Table below reports the status of `fyne-cross`'s supported cross compilation targets.
The supported architectures for the host are `amd64`and `arm64`. 

|                | amd64              | arm64 |
| -------------- | ------------------ | ----- |
| android/386    | :white_check_mark: |  N/A  |
| android/amd64  | :white_check_mark: |  N/A  |
| android/arm    | :white_check_mark: |  N/A  |
| android/arm64  | :white_check_mark: |  N/A  |
| darwin/amd64   | :white_check_mark: | :white_check_mark: |
| darwin/arm64   | :white_check_mark: | :white_check_mark: |
| freebsd/amd64  | :white_check_mark: | :white_check_mark: |
| freebsd/arm64  | :white_check_mark: | :white_check_mark: |
| linux/amd64    | :white_check_mark: | :white_check_mark: |
| linux/386      | :white_check_mark: | :white_check_mark: |
| linux/arm      | :white_check_mark: | :white_check_mark: |
| linux/arm64    | :white_check_mark: | :white_check_mark: |
| windows/amd64  | :white_check_mark: | :white_check_mark: |
| windows/386    | :white_check_mark: | :white_check_mark: |
| web            | :white_check_mark: | :white_check_mark: |

> Note: 
> - darwin images work against macOS SDK 11.3. Other SDKs are not supported.
> Other SDK versions could require a different min SDK version that can specified using the `--macosx-version-min` flag
> - android NDK is not available for linux/arm64, cross-compilation from arm64 hosts won't be supported in this initial release

## Building container with darwin sdk included

In some case, you might want to build your own container and include the darwin sdk in it. A solution to this is to use
`fyne-cross darwin-sdk-extract` command and copy the sdk you want in the subdirectory `darwin-with-sdk/sdk`. You can then use
the Makefile to build that new container. If you do not need to rebuild the `base` and `darwin` container, just reusing the
upstream version, you can do so by creating the appropriate `.`file that match the Makefile and the builder you are using to
create your image. For example: `touch .docker-base .docker-darwin` if you are using docker with the multi architecture Makefile.

## Contribute

- Fork and clone the repository
- Make and test your changes
- Open a pull request against the `develop` branch

### Contributors

See [contributors](https://github.com/fyne-io/fyne-cross-images/graphs/contributors) page
