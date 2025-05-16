#!/bin/bash

# Проверяем наличие .env файла
if [ ! -f .env ]; then
    echo "Файл .env не найден. Создаю из примера..."
    echo "WEATHER_API_KEY=your_openweather_api_key_here" > .env
    echo "PORT=8000" >> .env
    echo "Пожалуйста, отредактируйте файл .env и установите свой WEATHER_API_KEY"
    echo "Нажмите любую клавишу, чтобы продолжить..."
    read -n 1
fi

# Определяем режим работы
MODE=${1:-prod}

if [ "$MODE" == "dev" ]; then
    echo "Запуск в режиме разработки..."
    docker-compose -f stack.dev.yml up --build
else
    echo "Запуск в режиме продакшн..."
    docker-compose -f stack.yml up --build
fi