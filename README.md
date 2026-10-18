# Инструменты Granio

Сборка, упаковка и подготовка окружения приложений Granio.

**Основной исходник:** https://github.com/fyne-io/tools.
**Цель форка:** сопровождать инструменты, кросс-сборку, образы, мобильный мост
и среду разработки в одном репозитории. Истории, лицензии и авторские сведения
сохранены; дополнительные источники перечислены в [SOURCES.json](SOURCES.json).

| Компонент | Путь | Исходник |
| --- | --- | --- |
| Упаковщик | `cmd/granio` | [tools](https://github.com/fyne-io/tools) |
| Проверка окружения | `cmd/granio_setup`, библиотека `setup/pkg` | [setup](https://github.com/fyne-io/setup) |
| Кросс-сборка | `cross/cmd/granio_cross` | [fyne-cross](https://github.com/fyne-io/fyne-cross) |
| Рецепты образов | `cross-images` | [fyne-cross-images](https://github.com/fyne-io/fyne-cross-images) |
| Мобильный мост | `mobile-bridge` | [gomobile-bridge](https://github.com/fyne-io/gomobile-bridge) |
| Среда разработки | `studio` | [defyne](https://github.com/fyne-io/defyne) |

Go 1.27.1+. Кросс-сборка и мобильный мост сохраняют отдельные Go-модули:
зависимости Kubernetes и AWS не включаются в модуль обычного упаковщика.

```sh
go install github.com/Uzer007-admin/granio-tools/cmd/granio@{{TOOLS_VERSION}}
go install github.com/Uzer007-admin/granio-tools/cmd/granio_setup@{{TOOLS_VERSION}}
go install github.com/Uzer007-admin/granio-tools/cross/cmd/granio_cross@{{CROSS_VERSION}}
go install github.com/Uzer007-admin/granio-tools/studio@{{TOOLS_VERSION}}
```

Для локальной проверки каждого модуля задайте `GOWORK=off` и выполните
`go test -race -tags migrated_fynedo,ci ./...` в корне модуля.
Мобильный мост используется как `replace` для `golang.org/x/mobile`;
его новый адрес — `github.com/Uzer007-admin/granio-tools/mobile-bridge`.

Рецепты образов находятся в `cross-images`; сборка и публикация образов,
доступность платформенных SDK и приёмка мобильных приложений проверяются отдельно.
Флаги упаковщика, используемые существующими сценариями, сохраняются.
[Материалы исходного проекта](UPSTREAM.md).
