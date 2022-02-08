FROM node:latest

ENV LANG=C.UTF-8
ENV TZ=Asia/Tokyo

USER frontend

WORKDIR /frontend
COPY ./ /frontend

RUN npm install

ENTRYPOINT ["npm", "run", "dev"]

