from aiogram import Router, types, F, Bot
from aiogram.filters import Command
import json
import time
import logging
import hashlib
import hmac
import base64
from urllib.parse import urlencode

from ..keyboards.reply import get_welcome_keyboard, get_children_keyboard
from ..services.api import get_children
from ..texts.messages import WELCOME_MESSAGE, NO_CHILDREN_MESSAGE

router = Router()


def generate_telegram_hash(data: dict, bot_token: str) -> str:
    # Сортировка по ключам
    data_list = []
    for key in sorted(data.keys()):
        if key != "hash":
            value = data[key]
            data_list.append(f"{key}={value}")

    # Создание data-check-string
    data_check_string = "\n".join(data_list)

    # Создание secret_key
    secret_key = hmac.new(
        "WebAppData".encode(), bot_token.encode(), hashlib.sha256
    ).digest()

    # Создание хеша
    return hmac.new(secret_key, data_check_string.encode(), hashlib.sha256).hexdigest()


@router.message(Command("start"))
async def command_start(message: types.Message, webapp_url: str) -> None:
    await message.answer(
        text=WELCOME_MESSAGE,
        reply_markup=get_welcome_keyboard(webapp_url),
        parse_mode="HTML",
    )


@router.message(Command("kids"))
async def command_kids(
    message: types.Message, bot: Bot, api_url: str, kids_url: str
) -> None:
    try:
        auth_data = {
            "query_id": str(time.time()),
            "user": json.dumps(
                {
                    "id": message.from_user.id,
                    "first_name": message.from_user.first_name,
                    "last_name": message.from_user.last_name,
                    "username": message.from_user.username,
                    "language_code": message.from_user.language_code,
                }
            ),
            "auth_date": str(int(time.time())),
            "chat_type": "private",
        }

        auth_data["hash"] = generate_telegram_hash(auth_data, bot.token)

        children = await get_children(api_url, auth_data)
        if not children:
            await message.answer(text=NO_CHILDREN_MESSAGE)
            return

        await message.answer(
            text="Выберите ребенка:",
            reply_markup=get_children_keyboard(children, kids_url),
        )
    except Exception as e:
        logging.error(f"Error getting children: {e}")
        await message.answer(
            text=f"Произошла ошибка при получении списка детей: {str(e)}"
        )
