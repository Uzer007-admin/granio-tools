# granio-mobile-bridge

Компонент независимого проекта **Granio — Гранио**.

**Исходный репозиторий:** [gomobile-bridge](https://github.com/fyne-io/gomobile-bridge).
**Цель форка:** Мост мобильных библиотек и приложений Granio.

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

# Gomobile Bridge

The Fyne toolkit provides it's own mobile packaging tools which are based on gomobile.
Because of this applications that import gomobile (directly or through a dependency)
may not compile when targetting mobile devices. Or they may not work as expected if one
of the libraries expects gomobile to be loaded.

The gomobile-bridge tool solves this problem by simulating a gomobile API for fyne mobile
apps. This enabled libraries to access the same APIs and yet call the equivalent APIs in
the Fyne ecosystem. To make use of this simply add the following line to your `go.mod` file.

```
replace golang.org/x/mobile => github.com/fyne-io/gomobile-bridge v0.1.0
```

And then update your module info with:

```
go mod tidy
```

If you have a vendor directory it then you should update those files using the `vendor` command:

```
go mod vendor
```
