FROM node:8-alpine as frontend

COPY ./package.json ./package-lock.json /code/frontend/
WORKDIR /code/frontend
RUN npm install
# Now copy the rest of the codebase. This split of the steps ensures that the
# npm install stage will only be repeated if the package.json or
# package-lock.json files have changes.
COPY . /code/frontend
RUN npm run build

FROM nginx:1.13-alpine
COPY --from=frontend /code/frontend /usr/share/nginx/html
