---
title: Использование модулей Hugo
description: Как использовать модули Hugo.
categories: []
keywords: []
weight: 20
aliases: [/themes/usage/,/themes/installing/,/installing-and-using-themes/]
---

## Предварительные условия

{{% include "/_common/gomodules-info.md" %}}

## Инициализация нового модуля

Используйте `hugo mod init` для инициализации нового модуля Hugo. Если путь к модулю не удаётся угадать, его необходимо указать в качестве аргумента, например:

```sh
hugo mod init github.com/<your_user>/<your_project>
```

Также см. [CLI Doc](/commands/hugo_mod_init/).

## Использование модуля в качестве темы

Самый простой способ использования модуля в качестве темы — импортировать его в конфигурацию.

1. Инициализируйте систему модулей Hugo: `hugo mod init github.com/<your_user>/<your_project>`
1. Импортируйте тему:

    {{< code-toggle file=hugo >}}
    [module]
      [[module.imports]]
        path = "github.com/spf13/hyde"
    {{< /code-toggle >}}

## Обновление модулей

Модули будут загружены и добавлены, когда вы добавите их в качестве импорта в вашу конфигурацию. Смотрите [configure modules](/configuration/module/#imports).

Для обновления или управления версиями вы можете использовать `hugo mod get`.

Примеры:

### Обновление всех модулей

```sh
hugo mod get -u
```

### Обновить все модули рекурсивно

```sh
hugo mod get -u ./...
```

### Обновить конкретный модуль

```sh
hugo mod get -u github.com/gohugoio/myShortcodes
```

### Получить конкретную версию

```sh
hugo mod get github.com/gohugoio/myShortcodes@v1.0.7
```

Также см. [CLI Doc](/commands/hugo_mod_get/).

## Внесение и тестирование изменений в модуль

Одним из способов локальной разработки модуля, импортированного в проект, является добавление директивы replace в локальный каталог с исходным кодом в `go.mod`:

```sh
replace github.com/bep/hugotestmods/mypartials => /Users/bep/hugotestmods/mypartials
```

Если у вас запущен `hugo server`, конфигурация будет перезагружена, а `/Users/bep/hugotestmods/mypartials` будет добавлен в список наблюдения.

Вместо изменения файлов `go.mod` вы также можете использовать параметр конфигурации модулей [`replacements`](/configuration/module/#top-level-options).

## Печать графа зависиости

Используйте `hugo mod graph` из соответствующего каталога модулей, и он распечатает граф зависимости, включая авторство, замены модулей или отключенный статус.

Например:

```txt
hugo mod graph

github.com/bep/my-modular-site github.com/bep/hugotestmods/mymounts@v1.2.0
github.com/bep/my-modular-site github.com/bep/hugotestmods/mypartials@v1.0.7
github.com/bep/hugotestmods/mypartials@v1.0.7 github.com/bep/hugotestmods/myassets@v1.0.4
github.com/bep/hugotestmods/mypartials@v1.0.7 github.com/bep/hugotestmods/myv2@v1.0.0
DISABLED github.com/bep/my-modular-site github.com/spf13/hyde@v0.0.0-20190427180251-e36f5799b396
github.com/bep/my-modular-site github.com/bep/hugo-fresh@v1.0.1
github.com/bep/my-modular-site in-themesdir
```

Также см. [CLI Doc](/commands/hugo_mod_graph/).

## Распространяйте ваши модули (в оригинале - Vendor your modules)

`hugo mod vendor` запишет все зависимости модулей в каталог `_vendor`, который затем будет использоваться для всех последующих сборок.

Обратите внимание:

- Вы можете запустить `hugo mod vendor` на любом уровне в дереве модулей.
- Распространение не будет сохранять модули, хранящиеся в вашем каталоге `themes`.
- Большинство команд принимают флаг `--ignoreVendorPaths`, который затем не будет использовать вендорные модули в `_vendor` для путей модулей, соответствующих заданному шаблону [glob](g).

Также см. [CLI Doc](/commands/hugo_mod_vendor/).

## Аккуратно go.mod, go.sum

Запустите `hugo mod tidy`, чтобы удалить неиспользуемые записи в `go.mod` и `go.sum`.

Также см. [CLI Doc](/commands/hugo_mod_clean/).

## Очистить кэш модуля

Запустите `hugo mod clean`, чтобы удалить весь кэш модулей.

Обратите внимание, что вы также можете настроить кэш `modules` с `maxAge`. Смотрите [настроить кэши](/configuration/caches/).

Также см. [CLI Doc](/commands/hugo_mod_clean/).

## Рабочее пространство модуля

Поддержка рабочего пространства была добавлена ​​в [Go 1.18](https://go.dev/blog/get-familiar-with-workspaces), а Hugo получил надежную поддержку в версии `v0.109.0`.

Распространенным вариантом использования рабочего пространства является упрощение локальной разработки сайта с его тематическими модулями.

Рабочее пространство можно настроить в файле `*.work` и активировать с помощью параметра [module.workspace](/configuration/module/), который для этого использования обычно контролируется через переменную среды ОС `HUGO_MODULE_WORKSPACE`.

См. пример в файле [hugo.work](https://github.com/gohugoio/hugo/blob/master/docs/hugo.work) в репозитории Hugo Docs:

```text
go 1.20

use .
use ../gohugoioTheme
```

Используя директиву `use`, перечислите все модули, над которыми вы хотите работать, указав их относительное расположение. Как и в примере выше, рекомендуется всегда включать в список основной проект (`.`).

С этим вы можете запустить сервер Hugo с включенным рабочим пространством:

```sh
HUGO_MODULE_WORKSPACE=hugo.work hugo server --ignoreVendorPaths "**"
```

Флаг `--ignoreVendorPaths` добавлен выше, чтобы игнорировать любые зависимости от поставщиков внутри `_vendor`. Если вы не используете вендоринг, этот флаг вам не нужен. Но теперь сервер настроен на наблюдение за файлами и каталогами в рабочем пространстве, и вы можете видеть, как ваши локальные изменения перезагружаются.