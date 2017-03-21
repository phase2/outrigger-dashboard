FROM docker:17.03

COPY dist/outrigger-dashboard /outrigger-dashboard
COPY frontend/  /app

EXPOSE 80

CMD [ "/outrigger-dashboard" ]
