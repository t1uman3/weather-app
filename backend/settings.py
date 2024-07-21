import os
import dotenv


dotenv.load_dotenv('.env')

OPENWEATHER_API_KEY = os.environ["OPENWEATHER_API_KEY"]
OPENWEATHER_URL = os.environ["OPENWEATHER_URL"]
