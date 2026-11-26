# HTTP-Контракты

1) GET /app/v1/watch

* Watch отправляет запрос в Studio: GET /video

``` json
    Request: None

        Response: 
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

2) GET app/v1/watch/{channel_id}

* Watch отправляет запрос в Studio: GET /video/{channel_id}

``` json
    Request: None

    Response: [
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

3) GET app/v1/watch/{channel_id}/{video_id} 

* Watch отправляет запрос в Studio: GET /video/{channel_id}/{video_id}

``` json 
    Request: None

    Response: 
    {
    "id": 1,
    "channel_id": 1,
    "title": "Мой первый влог",
    "description": "Описание ролика",
    "views": 120,
    "likes": 15
    }
```

4) POST app/v1/watch/likes/{channel_id}/{video_id}

* Watch отправляет запрос в Studio: POST /video/like/{channel_id}/{video_id}

``` json
    Request: None

    Response: 
    {
        "success": true
    }  
``` 

5) POST app/v1/watch/subscribe/{channel_id}

* Watch отправляет запрос в Studio: POST /channel/subscribe/{channel_id}

``` json
    Request: None

    Response: 

    {
        "success": true
    }
```


6) POST /channel - Создать канал

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
7) POST /channel/subscribe/channel_id - подписаться на канал 

    ```sql 
        Request: None

        Response: 201 Created
    ```

8) DELETE /channel/channel_id - Удалить канал

    Request: None
    Response: 204 No content



9) POST /video/channel_id - создать видео на определенном канале
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
10) POST /video/like/channel_id/video_id - ставит лайк на видео 

    ``` sql 
        Request: None
        Response: 201 Created
    ```

11) GET /video/channel_id/video_id - отдает пользователю определенное видео на определенном канале

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


12) GET /video/channel_id - Получить все видео с канала

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

13) GET /video - Лентва рекомендаций([]Video, все видео со всех каналов с сортировкой по просмотрам)

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


14) DELETE /video/channel_id/video_id - удаляет определенное видео на определенном канале
    
    ``` sql
        Response: None
        Request: 204 No content
    ```

