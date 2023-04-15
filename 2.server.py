import asyncio
import json
import websockets

active_sockets = set()


async def handler_json(json_string):
    """
    Обработчик полученного json
    :param json_string:
    :return:
    """
    # Выводим полученный текст от клиента client.go
    print(json_string)
    # Отправляем ответ на server.go
    async with websockets.connect("ws://localhost:8899") as websocket:
        # json_string = {**json_string, "finished": 1}
        json_string = json.dumps(json_string)
        await websocket.send(json_string)


async def server(websocket, path):
    """
    Обработчик входящих соединений
    :param websocket:
    :param path:
    :return:
    """
    active_sockets.add(websocket)
    try:
        async for message in websocket:
            try:
                data = json.loads(message)
                await handler_json(data)
            except Exception as e:
                print(f"Ошибка при обработке сообщения: {e}")
    except websockets.exceptions.ConnectionClosedError:
        pass
    finally:
        active_sockets.remove(websocket)


async def main():
    """
    Главная функция
    :return:
    """
    # Вечный цикл работы сервера
    async with websockets.serve(server, "localhost", 8888):
        await asyncio.Future()


if __name__ == "__main__":
    asyncio.run(main())