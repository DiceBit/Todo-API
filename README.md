# ToDo-API

---

## Getting started

Для запуска программы, выполните команды:

```bash
git clone https://github.com/DiceBit/Todo-API.git
cd Todo-API
docker build -f ./build/Dockerfile -t todo-api-image . 
docker run --rm -p 8080:8080 todo-api-image 
```

---

## Endpoints

```text
[Get] /tasks
[Post] /tasks
[Put] /tasks/{id}
[Delete] /tasks/{id}
[Patch] /tasks/{id}/complete
```

### **[Get] /tasks**

Получение всех задач

#### Response

```json
[
  {
    "Title": "Task#1",
    "Description": "test1",
    "DueDate": "2024-11-15",
    "Overdue": true,
    "Completed": false
  },
  {
    "Title": "Task#2",
    "Description": "test2",
    "DueDate": "2023-11-15",
    "Overdue": true,
    "Completed": false
  }
]
```

### **[Post] /tasks**

Добавление задачи

#### Request

```json
{
  "title": "Task#3",
  "description": "test3",
  "dueDate": "2024-11-16"
}
```

#### Response

```json
{
  "Id": 3,
  "Title": "Task#3",
  "Description": "test3",
  "DueDate": "2024-11-16",
  "Overdue": false,
  "Completed": false
}
```

### **[Put] /tasks/{id}**

Обновление задачи

#### Request

```json
{
  "title": "Task#2",
  "description": "update task2",
  "dueDate": "2024-01-12"
}
```

#### Response

```json
{
  "Title": "Task#2",
  "Description": "update task2",
  "DueDate": "2024-01-12",
  "Overdue": true,
  "Completed": false
}
```

### **[Delete] /tasks/{id}**

Удаление задачи
#### Response
- **Status code:** `200 OK` **or** `404 Not Found`

### **[Patch] /tasks/{id}/complete**

Выполнение задачи

#### Request

```json
{
  "completed": true
}
```

#### Response

```json
{
  "Title": "Task#3",
  "Description": "test3",
  "DueDate": "2024-11-16",
  "Overdue": false,
  "Completed": true
}
```

## Форматы данных

- **id**: Уникальный идентификатор задачи (integer).
- **title**: Название задачи (string).
- **description**: Описание задачи (string).
- **dueDate**: Дата сдачи задачи (string)
- **overdue**: Статус просрочена ли задачи (bool)
- **completed**: Статус выполнена ли задача (bool)