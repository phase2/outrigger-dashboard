FROM docker:1.13

COPY dist/outrigger-dashboard /outrigger-dashboard
COPY frontend/  /app

EXPOSE 80

CMD [ "/outrigger-dashboard" ]
