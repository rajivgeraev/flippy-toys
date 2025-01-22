import asyncio
import logging
from aiohttp import web
from aiogram import Bot, Dispatcher
from aiogram.enums import ParseMode
from aiogram.client.default import DefaultBotProperties
from aiogram.webhook.aiohttp_server import SimpleRequestHandler
from aiogram import types

from .config import load_config
from .handlers import common


async def main():
    logging.basicConfig(
        level=logging.INFO,
        format="%(asctime)s - %(levelname)s - %(name)s - %(message)s",
    )

    config = load_config()

    default = DefaultBotProperties(parse_mode=ParseMode.HTML)
    bot = Bot(token=config.token, default=default)
    dp = Dispatcher()

    # Регистрация обработчиков
    dp.include_router(common.router)

    # Передача конфигурации в middleware
    dp["webapp_url"] = config.webapp_url
    dp["kids_url"] = config.kids_url
    dp["api_url"] = config.api_url

    # Установка команд бота
    await bot.set_my_commands(
        [
            types.BotCommand(command="start", description="Запустить бота"),
            types.BotCommand(command="kids", description="Режим Kids"),
        ]
    )

    # Создаем aiohttp приложение
    app = web.Application()

    # Добавляем health check endpoint
    async def health_check(request):
        return web.Response(text="OK")

    app.router.add_get("/health", health_check)

    # Запускаем веб-сервер и бота
    runner = web.AppRunner(app)
    await runner.setup()
    site = web.TCPSite(runner, "0.0.0.0", 8081)

    await site.start()
    await dp.start_polling(bot)


if __name__ == "__main__":
    asyncio.run(main())
