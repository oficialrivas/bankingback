package utils

import (
    "banking/models"
    "fmt"
)

// Simula el envío de notificaciones
func EnviarNotificacion(usuario models.User, mensaje string) {
    // Aquí iría la lógica para enviar la notificación real.
    // Por ahora, solo imprimiremos el mensaje en la consola para fines de demostración.
    fmt.Printf("Notificación para %s: %s\n", usuario.Nombre, mensaje)
}
