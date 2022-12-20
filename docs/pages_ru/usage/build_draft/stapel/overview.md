---
title: Обзор
permalink: usage/build_draft/stapel/overview.html
---

В werf встроен альтернативный синтаксис описания сборочных инструкций называемый Stapel, который даёт следующие возможности:

1. Удобство поддержки и параметризации комплексной конфигурации, возможность переиспользовать общие куски и генерировать конфигурацию однотипных образов за счет использования YAML-формата и шаблонизации.
2. Специальные инструкции для интеграции с git, позволяющие задействовать инкрементальную пересборку с учетом истории git-репозитория.
3. Наследование образов и импортирование файлов из образов (аналог multi-stage для Dockerfile).
4. Запуск произвольных сборочных инструкций, опции монтирования директорий и другие инструменты продвинутого уровня для сборки образов.
5. Более эффективная механика кеширования слоёв (сейчас в pre-альфа аналогичная схема поддерживается и для слоёв Dockerfile при сборке с Buildah).

<!-- TODO(staged-dockerfile): удалить 5 пункт как неактуальный -->

Сборка образов через сборщик Stapel предполагает описание сборочных инструкций в конфигурационном файле `werf.yaml`. Stapel поддерживается как для сборочного бекенда docker server (сборка через shell инструкции или ansible), так и для buildah (только shell-инструкции).

В данном разделе рассмотрено как описывать сборку образов с помощью сборщика Stapel, описание дополнительных возможностей и как ими пользоваться.