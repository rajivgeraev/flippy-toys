from dataclasses import dataclass
from os import getenv
from dotenv import load_dotenv


@dataclass
class Config:
    token: str
    webapp_url: str
    kids_url: str
    api_url: str


def load_config() -> Config:
    load_dotenv()

    return Config(
        token=getenv("BOT_TOKEN"),
        webapp_url=getenv("WEBAPP_URL", "https://app.flippy.toys"),
        kids_url=getenv("KIDS_URL", "https://kids.flippy.toys"),
        api_url=getenv("API_URL", "https://api.flippy.toys"),
    )
