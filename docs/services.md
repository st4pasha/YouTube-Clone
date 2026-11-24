# Сервисы

1) Сервис `watch`, где мы просто смотрим видео

    Сервис `watch` будет посылать запрос к сервису `studio`, чтобы получить определенные видео или же все

* GET /watch - лента рекомендаций 

    Request: None

    Response: 200 OK

    ```sql
    [
    {
        "id": 2,
        "channel_id": 1,
        "title": "Самое популярное видео",
        "description": "Описание",
        "views": 1500000,
        "likes": 85000
    },
    {
        "id": 1,
        "channel_id": 1,
        "title": "Мой первый влог",
        "description": "Описание ролика",
        "views": 120,
        "likes": 15
    }
    ]
    ```

* GET /watch/channel_id - посмотреть все видео на канале

    ```sql
    Request: None

    Response: 200 OK

    [
    {
        "id": 1,
        "channel_id": 1,
        "title": "Мой первый влог",
        "description": "Описание ролика",
        "views": 120,
        "likes": 15
    }
    ]
    ```

* GET /watch/channel_id/video_id - посмотреть определенное видео на канале

    ```sql
    Request: None

    Response: 200 OK

    {
    "id": 1,
    "channel_id": 1,
    "title": "Мой первый влог",
    "description": "Описание ролика",
    "views": 120,
    "likes": 15
    }


* POST /watch/likes/channel_id/video_id - поставить лайк на видео

    ``` sql
    Request: None

    Response: 201 Created

    {
        "success": true
    }  
    ```
* POST /watch/subscribe/channel_id - подписаться на канал

    ``` sql
        Request: None

        Response: 201 Created
        
        {
            "success": true
        }

    ```

В сервисе watch ничего не хранится, все получается путем отправкеи запроса в сервис studio

2) Сервис studio, на нем будет управление каналами, добавление видео, хранение видео
    
* Фичи:

    1) Фича `channel`:

* POST /channel - Создать канал

        ```sql
        Request:
        {
            "channel_name": "MisterBist"
        }

        Response: 200 OK
        {
            "id": 1,
            "channel_name": "MisterBist",
            "followers":, 50000000
        }
        ```
* POST /channel/subscribe/channel_id - подписаться на канал 

    ```sql 
        Request: None

        Response: 201 Created
    ```

* DELETE /channel/channel_id - Удалить канал

    Request: None
    Response: 204 No content

    2) Фича `video`:

* POST /video/channel_id - создать видео на определенном канале
    ```sql
    Request 
    {
        "title": "title",
        "description": "description"
    }

    Response: 201 Created {
        "id": 1,
        "title": "title",
        "description": "description",
        "views":, 0,
        "likes":, 0
    }
    ```
* POST /video/like/channel_id/video_id - ставит лайк на видео 

    ``` sql 
        Request: None
        Response: 201 Created
    ```

* GET /video/channel_id/video_id - отдает пользователю определенное видео на определенном канале

    ```sql
    Request: None

    Response: 200 OK {
    "id": 1,
    "channel_id": 1,
    "title": "title",
    "description": "description",
    "views": 0,
    "likes": 0
    }
    ```


* GET /video/channel_id - Получить все видео с канала

    ```sql
    Response: 200 OK
    [
    {
        "id": 1,
        "channel_id": 1,
        "title": "title",
        "description": "description",
        "views": 0,
        "likes": 0
    }
    ]
    ```

* GET /video - Лентва рекомендаций([]Video, все видео со всех каналов с сортировкой по просмотрам)

    ``` sql
    Request: None

    Response: 200 OK
    [
    {
        "id": 2,
        "channel_id": 1,
        "title": "Самое популярное видео",
        "description": "Описание",
        "views": 1500000,
        "likes": 85000
    },
    {
        "id": 1,
        "channel_id": 1,
        "title": "title",
        "description": "description",
        "views": 0,
        "likes": 0
    }
    ]
    ```


* DELETE /video/channel_id/video_id - удаляет определенное видео на определенном канале
    
    ``` sql
        Response: None
        Request: 204 No content
    ```

