# cotizacion-telegram

Servicio Go que lee la SQLite compartida y publica en Telegram usando las
imágenes creadas previamente por `cotizacion_social`.

## Flujo

Este proyecto se ejecuta después de:

1. `cotizaciones`, que actualiza la DB y genera el JSON.
2. `cotizacion_social`, que genera las imágenes.
3. `cotizacion-telegram`, que lee la DB, usa las imágenes y envía o edita los
   dos mensajes de Telegram.

## Rutas de producción

- SQLite: `/opt/osbo/datausd`
- Imagen USD/USDT: `/opt/osbo/cotiza/usdt.jpg`
- Imagen resto: `/opt/osbo/cotiza/resto.jpg`
- Lock: `/opt/osbo/cotizaciones.lock`

El servicio conserva en la tabla `config` los IDs `messageid` y `messageidusd`
y los umbrales `umbral` y `umbral_referencial`. También conserva la detección
de spikes, reintentos seguros, fallback de edición foto/texto/caption y fallback
de foto a mensaje de texto cuando una imagen no está disponible.

## Desarrollo

```bash
go test ./...
go build -o cotizacion-telegram .
```

Requiere Go 1.27.3 y `TELEGRAM_BOT_TOKEN`.

## CI/CD

El workflow de GitHub Actions compila un binario Linux estático y, al hacer
push a `main`, lo publica en el registry local del servidor como
`localhost:32000/cotizacion-telegram:latest`. Después reaplica el manifiesto
del CronJob de `bot-dockers`.
