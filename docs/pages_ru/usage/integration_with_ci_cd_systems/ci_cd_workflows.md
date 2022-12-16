---
title: Рабочие процессы CI/CD
permalink: usage/integration_with_ci_cd_systems/ci_cd_workflows.html
---

## Готовые конфигурации workflow

Мы предлагаем пользователю на выбор несколько готовых конфигураций workflow для проекта. Эти конфигурации составлены из приведенных выше блоков workflow. В документации эти готовые конфигурации могут называться также стратегиями workflow.

Конкретные конфиги по каждой из конфигураций можно найти в инструкциях по конкретной CI/CD системе. Например, Gitlab CI: ссылка, Github Actions: ссылка.

### №1 Fast and Furious

Конфигурация рекомендована в качестве наиболее соответствующей канонам CI/CD, которую можно реализовать с помощью werf.

В данной конфигурации может быть произвольное число production-like окружений, как то: testing, staging, development, qa, и т.д.

| **Окружение** | **Блок workflow** |
| :--- | :--- |
| Production | [Выкат на production из master автоматически](#выкат-на-production-из-master-автоматически) + откат через revert |
| Staging / Testing / Development / QA | [Выкат на production-like из pull request по кнопке](#выкат-на-production-like-из-pull-request-по-кнопке) |
| Review | [Выкат на review из pull request автоматически после ручной активации](#выкат-на-review-из-pull-request-автоматически-после-ручной-активации) |

### №2 Push the button

| **Окружение** | **Блок workflow** |
| :--- | :--- |
| Production | [Выкат на production из master по кнопке](#выкат-на-production-из-master-по-кнопке) |
| Staging | [Выкат на staging из master автоматически](#выкат-на-staging-из-master-автоматически) |
| Testing / Development / QA | [Выкат на production-like из ветки автоматически](#выкат-на-production-like-из-ветки-автоматически) |
| Review | [Выкат на review из pull request по кнопке](#выкат-на-review-из-pull-request-по-кнопке) |

### №3 Tag everything

Не все проекты сходу готовы к внедрению CI/CD. В таких проектах используется более классический метод создания релизов только после активной фазы разработки. Переход к CI/CD в таких проектах требует усилий по преодолению привычных вещей и переосмысления как от разработчиков, так и от devops. Поэтому для таких проектов мы предлагаем и классическую конфигурацию, которая также рекомендована для werf в случае невозможности использования [fast & furious](#1-fast-and-furious).

| **Окружение** | **Блок workflow** |
| :--- | :--- |
| Production | [Выкат на production из тега автоматически](#выкат-на-production-из-тега-автоматически) |
| Staging | [Выкат на staging из master автоматически](#выкат-на-staging-из-master-автоматически) или [выкат на staging из master по кнопке](#выкат-на-staging-из-master-по-кнопке) |
| Review | [Выкат на review из pull request автоматически после ручной активации](#выкат-на-review-из-pull-request-автоматически-после-ручной-активации) |

### №4 Branch, branch, branch

Управляем выкатом прямо через git с использованием веток и процедур git merge, rebase и push-force. Через создание определённых имен веток получаем автоматический выкат на review окружения.

Рекомендуем для тех, кто хочет управлять CI/CD полностью из git. Отметим, что подход также является соответствующим канонам CI/CD, как и fast & furious.

| **Окружение** | **Блок workflow** |
| :--- | :--- |
| Production | [Выкат на production из master автоматически](#выкат-на-production-из-master-автоматически) + откат через revert |
| Staging | [Выкат на staging из master автоматически](#выкат-на-staging-из-master-автоматически) |
| Review | [Выкат на review из ветки по шаблону автоматически](#выкат-на-review-из-ветки-по-шаблону-автоматически) |

## Составляющие workflow для отдельных окружений

Далее рассмотрим различные варианты выката production и других окружений в связке с git. Каждый пункт определяет строительный блок, который можно использовать для работы с определённым окружением. Мы будем называть такой строительный блок блоком workflow. Из блоков workflow в дальнейшем можно собрать свой workflow или взять готовую конфигурацию (см. далее [готовые конфигурации workflow](#готовые-конфигурации-workflow)).

Review окружения создаются и удаляются динамически по требованию разработчиков. С этим связаны особенности выката в эти окружения. В разделах, связанных с review, будет описано не только как создать review-окружение, но и как его удалить.

### Выкат на production из master автоматически

Merge или коммит в ветку master вызывает pipeline выката непосредственно на production.

Состояние ветки в любой момент времени отражает состояние окружения. Поэтому данный вариант является соответствующим подходу true CI/CD.

Варианты отката:
- Рекомендованный: откат через реверт коммита в ветке master. В этом случае поддерживается состояние ветки в синхронизированном с окружением состоянии, поэтому это предпочтительный вариант для сохранения целостности схемы.
- Средствами CI/CD системы, повторный [ручной вызов pipeline](#варианты-ручного-запуска-pipeline) на старом коммите (например, в Gitlab CI кнопка "откатить" по факту выполняет именно эти шаги).

### Выкат на production из master по кнопке

Pipeline выката в production может быть запущен вручную только на коммите из ветки master. Запуск pipeline производится средствами CI/CD системы [вручную](#варианты-ручного-запуска-pipeline): кнопка в CI/CD системе или вызов API.

Варианты отката:
- Рекомендованный: средствами CI/CD системы, повторный [ручной вызов pipeline](#варианты-ручного-запуска-pipeline) на старом коммите (например, в Gitlab CI кнопка "откатить" по факту выполняет именно эти шаги).
- Реверт коммита в ветке master, затем запуск pipeline средствами CI/CD системы [вручную](#варианты-ручного-запуска-pipeline): кнопка в CI/CD системе или вызов API. В данном случае вариант не рекомендован, т.к. состояние мастера не всегда соответствует состоянию окружения (в отличие от варианта master-автоматом), поэтому создавать лишний revert не имеет большого смысла именно для задачи отката.

### Выкат на production из тега автоматически

Создание нового тега автоматически вызывает pipeline выката на production-окружение из коммита, связанного с этим тегом.

Варианты отката:
- Рекомендованный: средствами CI/CD системы, повторный [ручной вызов pipeline](#варианты-ручного-запуска-pipeline) на старом теге.
- Создание нового тега на старый коммит, далее автоматический вызов pipeline выката для нового тега. Не предпочтительный вариант, т.к. плодятся лишние теги.

### Выкат на production из тега по кнопке

Pipeline выката в production-окружение может быть вызван только на существующем теге в git. Запуск pipeline производится средствами CI/CD системы [вручную](#варианты-ручного-запуска-pipeline): кнопка в CI/CD системе или вызов API.

Варианты отката:
- Средствами CI/CD системы, повторный [ручной вызов pipeline](#варианты-ручного-запуска-pipeline) на старом теге.

### Выкат на production из ветки автоматически

Merge или коммит в специальную ветку вызывает pipeline выката непосредственно на production (вариант похож на (master-автоматически)(#master-автоматически), но используется отдельная ветка).

Состояние специальной ветки в любой момент времени отражает состояние окружения. Поэтому данный вариант является соответствующим подходу true CI/CD.

Варианты отката:
- Рекомендованный: откат через реверт коммита в ветке. В этом случае поддерживается состояние ветки в синхронизированном с окружением состоянии, поэтому это предпочтительный вариант для сохранения целостности схемы.
- Средствами CI/CD системы, повторный [ручной вызов pipeline](#варианты-ручного-запуска-pipeline) на старом коммите (например, в Gitlab CI кнопка "откатить" по факту выполняет именно эти шаги).
- Реверт коммита в master, затем fast-forward merge в специальную ветку.
- Удаление коммита в специальной ветке (через удаление коммита в git, затем процедура git push-force).

### Выкат на production из ветки по кнопке

Pipeline выката в production может быть запущен вручную только на коммите из специальной ветки. Запуск pipeline производится средствами CI/CD системы [вручную](#варианты-ручного-запуска-pipeline): кнопка в CI/CD системе или вызов API.

Варианты отката:
- Рекомендованный: средствами CI/CD системы, повторный [ручной вызов pipeline](#варианты-ручного-запуска-pipeline) на старом коммите (например, в Gitlab CI кнопка "откатить" по факту выполняет именно эти шаги).
- Реверт коммита в ветке, затем запуск pipeline средствами CI/CD системы [вручную](#варианты-ручного-запуска-pipeline): кнопка в CI/CD системе или вызов API. В данном случае вариант не рекомендован, т.к. состояние мастера не всегда соответствует состоянию окружения (в отличие от варианта master-автоматом), поэтому создавать лишний revert не имеет большого смысла именно для задачи отката.

### Выкат на production-like из pull request по кнопке

Pipeline выката в production может быть запущен на любом коммите в pull request. Запуск pipeline производится средствами CI/CD системы [вручную](#варианты-ручного-запуска-pipeline): кнопка в CI/CD системе или вызов API.

Варианты отката:
- Средствами CI/CD системы, повторный [ручной вызов pipeline](#варианты-ручного-запуска-pipeline) на старом коммите (например, в Gitlab CI кнопка "откатить" по факту выполняет именно эти шаги).

### Выкат на staging из master автоматически

Merge или коммит в ветку master вызывает pipeline выката непосредственно на staging окружение.

Состояние ветки в любой момент времени отражает состояние окружения. Поэтому данный вариант является соответствующим подходу true CI/CD.

Варианты отката:
- Рекомендованный: откат через реверт коммита в ветке master. В этом случае поддерживается состояние ветки в синхронизированном с окружением состоянии, поэтому это предпочтительный вариант для сохранения целостности схемы.
- Средствами CI/CD системы, повторный [ручной вызов pipeline](#варианты-ручного-запуска-pipeline) на старом коммите (например, в Gitlab CI кнопка "откатить" по факту выполняет именно эти шаги).

### Выкат на staging из master по кнопке

Pipeline выката в staging может быть запущен вручную только на коммите из ветки master. Запуск pipeline производится средствами CI/CD системы [вручную](#варианты-ручного-запуска-pipeline): кнопка в CI/CD системе или вызов API.

Варианты отката:
- Рекомендованный: средствами CI/CD системы, повторный [ручной вызов pipeline](#варианты-ручного-запуска-pipeline) на старом коммите (например, в Gitlab CI кнопка "откатить" по факту выполняет именно эти шаги).
- Реверт коммита в ветке master, затем запуск pipeline средствами CI/CD системы [вручную](#варианты-ручного-запуска-pipeline): кнопка в CI/CD системе или вызов API. В данном случае вариант не рекомендован, т.к. состояние мастера не всегда соответствует состоянию окружения (в отличие от варианта Выкат на staging из master автоматически), поэтому создавать лишний revert не имеет большого смысла именно для задачи отката.

### Выкат на production-like из ветки автоматически

Merge или коммит в специальную ветку вызывает pipeline выката непосредственно на production-like окружение (вариант похож на (master-автоматически)(#master-автоматически), но используется отдельная ветка). Для каждого конкретного production-like окружения, как то: staging или testing — используется отдельная ветка.

Состояние специальной ветки в любой момент времени отражает состояние окружения. Поэтому данный вариант является соответствующим подходу true CI/CD.

Варианты отката:
- Рекомендованный: откат через реверт коммита в ветке. В этом случае поддерживается состояние ветки в синхронизированном с окружением состоянии, поэтому это предпочтительный вариант для сохранения целостности схемы.
- Средствами CI/CD системы, повторный [ручной вызов pipeline](#варианты-ручного-запуска-pipeline) на старом коммите (например, в Gitlab CI кнопка "откатить" по факту выполняет именно эти шаги).
- Реверт коммита в master, затем fast-forward merge в специальную ветку.
- Удаление коммита в специальной ветке (через удаление коммита в git, затем процедура git push-force).

### Выкат на production-like из ветки по кнопке

Pipeline выката в production-like окружение может быть запущен вручную только на коммите из специальной ветки. Запуск pipeline производится средствами CI/CD системы [вручную](#варианты-ручного-запуска-pipeline): кнопка в CI/CD системе или вызов API.

Варианты отката:
- Рекомендованный: средствами CI/CD системы, повторный [ручной вызов pipeline](#варианты-ручного-запуска-pipeline) на старом коммите (например, в Gitlab CI кнопка "откатить" по факту выполняет именно эти шаги).
- Реверт коммита в ветке, затем запуск pipeline средствами CI/CD системы [вручную](#варианты-ручного-запуска-pipeline): кнопка в CI/CD системе или вызов API. В данном случае вариант не рекомендован, т.к. состояние ветки не всегда соответствует состоянию окружения (в отличие от варианта Выкат на production-like из ветки автоматически), поэтому создавать лишний revert не имеет большого смысла именно для задачи отката.

### Выкат на review из pull request автоматически

Создание pull request автоматически вызывает выкат в отдельное review окружение. Название этого окружения связано с именем ветки. Дальнейшие коммиты в ветку, связанную с pull request автоматически вызывают выкат в review окружение.

Варианты отката:
- Рекомендованный: откат через реверт коммита в ветке. В этом случае поддерживается состояние ветки в синхронизированном с окружением состоянии, поэтому это предпочтительный вариант для сохранения целостности схемы.
- Средствами CI/CD системы, повторный [ручной вызов pipeline](#варианты-ручного-запуска-pipeline) на старом коммите (например, в Gitlab CI кнопка "откатить" по факту выполняет именно эти шаги).

Удаление review-окружения:
- По закрытию или принятию PR.
- Автоматически по истечению time-to-live с последнего выката на данное окружение (другими словами, при отсутствии активности в данном окружении).

### Выкат на review из ветки по шаблону автоматически

Создание pull request для ветки, подходящей под определённый паттерн автоматически вызывает выкат в отдельное review окружение. Название этого окружения связано с именем ветки. Дальнейшие коммиты в ветку, связанную с pull request автоматически вызывают выкат в review окружение.

Например, для паттерна `review_*` создание pull request для ветки `review_myfeature1` вызовет автоматическое создание соответствующего review окружения.

Варианты отката:
- Рекомендованный: откат через реверт коммита в ветке. В этом случае поддерживается состояние ветки в синхронизированном с окружением состоянии, поэтому это предпочтительный вариант для сохранения целостности схемы.
- Средствами CI/CD системы, повторный [ручной вызов pipeline](#варианты-ручного-запуска-pipeline) на старом коммите (например, в Gitlab CI кнопка "откатить" по факту выполняет именно эти шаги).

Удаление review-окружения:
- По закрытию или принятию PR.
- Автоматически по истечению time-to-live с последнего выката на данное окружение (другими словами, при отсутствии активности в данном окружении).

### Выкат на review из pull request по кнопке

Pipeline выката в review-окружение может быть запущен вручную только на коммите из ветки соответствующей этому окружению. Название этого окружения связано с именем ветки. Запуск pipeline производится средствами CI/CD системы [вручную](#варианты-ручного-запуска-pipeline): кнопка в CI/CD системе, повесить label или вызов API.

Варианты отката:
- Рекомендованный: средствами CI/CD системы, повторный [ручной вызов pipeline](#варианты-ручного-запуска-pipeline) на старом коммите (например, в Gitlab CI кнопка "откатить" по факту выполняет именно эти шаги).
- Реверт коммита в ветке, затем запуск pipeline средствами CI/CD системы [вручную](#варианты-ручного-запуска-pipeline): кнопка в CI/CD системе или вызов API. В данном случае вариант не рекомендован, т.к. состояние ветки не всегда соответствует состоянию окружения (в отличие от вариантов "автоматом для pull-request" и "автоматом для pull-request по паттерну"), поэтому создавать лишний revert не имеет большого смысла именно для задачи отката.

Удаление review-окружения:
- По закрытию или принятию PR.
- Автоматически по истечению time-to-live с последнего выката на данное окружение (другими словами, при отсутствии активности в данном окружении).

### Выкат на review из pull request автоматически после ручной активации

Review-окружение для pull request создаётся после его ручной активации средствами CI/CD системы. С этого момента любой коммит в ветку, связанную с pull request, вызывает автоматический выкат на review окружение. После работы с review его можно деактивировать вручную средствами CI/CD системы.

Pipeline выката в review-окружение может быть запущен только на коммите из ветки соответствующей этому окружению. Название этого окружения связано с именем ветки. Запуск pipeline для активации review окружения производится средствами CI/CD системы [вручную](#варианты-ручного-запуска-pipeline): кнопка в CI/CD системе, повесить label или вызов API.

Варианты отката:
- Рекомендованный: откат через реверт коммита в ветке. В этом случае поддерживается состояние ветки в синхронизированном с окружением состоянии, поэтому это предпочтительный вариант для сохранения целостности схемы.
- Средствами CI/CD системы, повторный [ручной вызов pipeline](#варианты-ручного-запуска-pipeline) на старом коммите (например, в Gitlab CI кнопка "откатить" по факту выполняет именно эти шаги).

Удаление review-окружения:
- Запуск pipeline для деактивации review окружения средствами CI/CD системы [вручную](#варианты-ручного-запуска-pipeline): кнопка в CI/CD системе, снять ранее повешенный label или вызов API.
- По закрытию или принятию PR.
- Автоматически по истечению time-to-live с последнего выката на данное окружение (другими словами, при отсутствии активности в данном окружении).