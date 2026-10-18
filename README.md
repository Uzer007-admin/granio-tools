# granio-setup

Компонент независимого проекта **Granio — Гранио**.

**Исходный репозиторий:** [setup](https://github.com/fyne-io/setup).
**Цель форка:** Подготовка окружения разработчика Granio.

История изменений, лицензии и сведения об исходных авторах сохранены.
Go-модули переведены на собственные пути `github.com/Uzer007-admin/…`;
минимальная версия Go — **1.27.1**. Ядро и общие API подключаются из
[`granio/v2`](https://github.com/Uzer007-admin/granio).

Состояние переноса и результаты проверок приведены в
[реестре экосистемы](https://github.com/Uzer007-admin/granio/blob/dev/granio/docs/ECOSYSTEM.md).
Наличие форка само по себе не означает проверку всех платформ и сценариев.

---

## Документация исходного проекта

Команды установки и ссылки в сохранённом тексте ниже относятся к исходнику.
Для Granio используются новые пути модулей, перечисленные выше.

# Fyne Setup

## Synopsis

The repository hosts a graphical tool `setup` which aims
to help developers get their environment set up.

You can run this graphical interface to see if your computer is ready to go!

You should see something like the following screenshot:

![Setup screenshot](img/screenshot.png)

If your setup is not ready then it will give hints about how
to complete the configuration.

## Binaries

You may obtain binaries from `setup` tool from this [download page](https://geoffrey-artefacts.fynelabs.com/github/andydotxyz/fyne-io/setup/latest/).

## Installation

If you prefer to build from source, follow the instructions below.

### Prerequisites

Please make sure you have go, gcc and graphics library header files installed on your system.
You may consult the section [Prerequisites](https://docs.fyne.io/started/) in the documentation.

### Building and running

Use the commands below to build and run the tool from sources:

```
go install fyne.io/setup@latest
$(go env GOPATH)/bin/setup
```
