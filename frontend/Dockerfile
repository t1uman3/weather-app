# Используем официальный образ Node
FROM node:16 AS build

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файл package.json и package-lock.json
COPY package*.json ./

# Устанавливаем зависимости системы
RUN npm install

# Копируем остальные файлы проекта
COPY . .

# Запускаем проект
RUN npm run build

FROM nginx:alpine

COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 3000

CMD ["nginx", "-g", "daemon off;"]