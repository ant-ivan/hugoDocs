---
title: Введение 
description: Краткое введение в модули Hugo.
categories: []
keywords: []
weight: 10
---

Hugo uses modules as its fundamental organizational units. A module can be a full Hugo project or a smaller, reusable piece providing one or more of Hugo's seven component types: static files, content, layouts, data, assets, internationalization (i18n) resources, and archetypes.

Hugo использует модули в качестве основных организационных единиц. Модуль может быть полным проектом Hugo или меньшей, повторно используемой частью, предоставляющей один или несколько из семи типов компонентов Hugo: статические файлы, контент, макеты, данные, активы, ресурсы интернационализации (i18n) и архетипы.

Modules are combinable in any arrangement, and external directories (including those from non-Hugo projects) can be mounted, effectively creating a single, unified file system.

Модули можно комбинировать в любом порядке, а внешние каталоги (включая каталоги из проектов, не относящихся к Hugo) можно монтировать, эффективно создавая единую унифицированную файловую систему.

Some example projects:

Некоторые примеры проектов:

[https://github.com/bep/docuapi](https://github.com/bep/docuapi)
: A theme that has been ported to Hugo Modules while testing this feature. It is a good example of a non-Hugo-project mounted into Hugo's directory structure. It even shows a JS Bundler implementation in regular Go templates.

: тема, которая была перенесена в Hugo Modules во время тестирования этой функции. Это хороший пример не-Hugo-проекта, смонтированного в структуру каталогов Hugo. Он даже показывает реализацию JS Bundler в обычных шаблонах Go.

[https://github.com/bep/my-modular-site](https://github.com/bep/my-modular-site)
: A simple site used for testing.

: Простой сайт, используемый для тестирования.