import asyncio
import json
from config import HOST, ROOT
from tracer import Model
import websockets
import db


connection = db.connect_db(HOST)
# db.creator(connection)
# db.deleter(connection)
data, err = db.fetch_data(connection.cursor())
rat = Model(ROOT, data=data)


active_sockets = set()


async def handler_json(json_string):
    """
    Обработчик полученного json
    :param json_string:
    :return:
    """
    # Выводим полученный текст от клиента client.go
    print(json_string)
    req = [int(i) for i in json_string["Details"]]
    request = {"summ": sum(req), "data": req}
    prediction = rat.search(request)

    print(prediction)
    # Отправляем ответ на server.go
    async with websockets.connect("ws://localhost:8899") as websocket:
        json_string = {**json_string, "Recommend": prediction}
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