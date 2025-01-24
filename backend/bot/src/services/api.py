import aiohttp
from typing import List
import json
from urllib.parse import urlencode
from ..models import Child


async def get_children(api_url: str, init_data: dict) -> List[Child]:
    async with aiohttp.ClientSession() as session:
        headers = {
            "Content-Type": "application/json",
            "X-Telegram-Init-Data": urlencode(init_data),
        }
        async with session.get(
            f"{api_url}/api/v1/children", headers=headers
        ) as response:
            data = await response.json()
            return [Child(**child) for child in data]
