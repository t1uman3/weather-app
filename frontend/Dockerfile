# Используем официальный образ Node с поддержкой Alpine Linux
FROM node:18-alpine AS build

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы для установки зависимостей
COPY package*.json ./

# Устанавливаем зависимости с кэшированием
RUN npm ci --silent

# Копируем остальные файлы проекта
COPY . .

# Собираем проект
RUN npm run build

# Этап production с Nginx
FROM nginx:stable-alpine

# Копируем конфигурацию Nginx
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Копируем собранные файлы из этапа сборки
COPY --from=build /app/dist /usr/share/nginx/html

# Открываем порт
EXPOSE 80

# Запускаем Nginx
CMD ["nginx", "-g", "daemon off;"]