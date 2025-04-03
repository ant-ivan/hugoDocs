---
title: Введение 
description: Краткое введение в модули Hugo.
categories: []
keywords: []
weight: 10
---

Hugo uses modules as its fundamental organizational units. A module can be a full Hugo project or a smaller, reusable piece providing one or more of Hugo's seven component types: static files, content, layouts, data, assets, internationalization (i18n) resources, and archetypes.

Hugo использует модули в качестве основных организационных единиц. Модуль может быть полным проектом Hugo или меньшей, повторно используемой частью, предоставляющей один или несколько из семи типов компонентов Hugo: статические файлы _(static files)_, контент _(content)_, макеты _(layouts)_, данные _(data)_, активы _(assets)_, ресурсы интернационализации _(internationalization (i18n))_ и архетипы _(archetypes)_.

Модули можно комбинировать в любом порядке, а внешние каталоги (включая каталоги из проектов, не относящихся к Hugo) можно монтировать, эффективно создавая единую унифицированную файловую систему.

Некоторые примеры проектов:

[https://github.com/bep/docuapi](https://github.com/bep/docuapi)
: Тема, которая была перенесена в Hugo Modules во время тестирования этой функции. Это хороший пример не-Hugo-проекта, смонтированного в структуру каталогов Hugo. Он даже показывает реализацию JS Bundler в обычных шаблонах Go.

[https://github.com/bep/my-modular-site](https://github.com/bep/my-modular-site)
: Простой сайт, используемый для тестирования.