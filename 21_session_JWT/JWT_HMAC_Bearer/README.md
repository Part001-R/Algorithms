### Пример работы.
Пользователь проходит регистрацию. Создаётся хеш токена сессии из имени пользователя и пароля.
{
    "user_name": "Alex",
    "user_pwd": "AAA",
    "user_pwd_repeat": "AAA"
}

Пользователь проходит аутентификацию. Клиенту, возвращается JWT. `Aythorization: Bearer <JWT>`
{
    "user_name": "Alex",
    "user_pwd": "AAA"
}

Пользователь запрашивает данные.

### Передача сессии.
JWT передаётся в заголовке Authorization, с префиксом Bearer.

Пример: Aythorization: Bearer <JWT>
