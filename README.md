# CLI Todo Application

### Консольное приложение для управления задачами, написанное на Go. Поддерживает категории, приоритеты, статусы и временные метки.

## Возможности
1. Создание, редактирование и удаление задач
2. Организация по категориям
3. Управление приоритетами (high/medium/low/none)
4. Статусы выполнения (todo/inprogress/done)
5. Дедлайн задач
6. Статистика по задачам
7. Фильтрация

## Установка
Клонируйте репозиторий:

`git clone https://github.com/ZaharBorisenko/Cli-App.git
cd Cli-App`

Компиляция приложения:

`go build -o todo.exe main.go` -> `.\todo.exe -list`

## Все команды приложения

* `.\todo.exe -add "Название задачи"`
* `.\todo.exe -edit "id:новое_название"`
* `.\todo.exe -addDesc "id:описание"`
* `.\todo.exe -toggle id`
* `.\todo.exe -del id`
* `.\todo.exe -list`
* `.\todo.exe -stats`
* `.\todo.exe -addCategory "Название категории"`
* `.\todo.exe -setCategory "id:категория"`
* .`\todo.exe -listByCat "категория"`
* `.\todo.exe -listCats`
* `.\todo.exe -setPriority "id:приоритет"` - **high / low / medium**
* `.\todo.exe -listByPriority "приоритет"`  **high / low / medium**
* `.\todo.exe -colors=false`  **true / false**
* `.\todo.exe -setStatus "id:статус"`  **high / low / medium**
* `.\todo.exe -listByStatus "статус"` **todo / in_progress/ done**
* `.\todo.exe -listDone`
* .`\todo.exe -listActive`
* `.\todo.exe -setTime "id:дата"` **format - 02/01/2025**