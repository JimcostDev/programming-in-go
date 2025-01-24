# 📖 Resumen: Concurrencia y Paralelismo en Go

Go es un lenguaje diseñado para aprovechar al máximo la **concurrencia** y el **paralelismo**, permitiendo a los desarrolladores manejar múltiples tareas de forma eficiente. Este documento resume conceptos clave, herramientas y ejemplos relacionados con este tema.

---

## **📊 Tabla Comparativa: Concurrencia vs Paralelismo**

| **Característica**         | **Concurrencia**                                               | **Paralelismo**                          |
|-----------------------------|---------------------------------------------------------------|------------------------------------------|
| **Definición**              | Varias tareas progresan de manera intercalada.                | Varias tareas se ejecutan al mismo tiempo. |
| **Requiere Multinúcleo**    | No necesariamente.                                             | Sí.                                      |
| **Ejemplo**                | Manejar múltiples solicitudes web mientras esperas respuestas. | Ejecutar cálculos intensivos en núcleos separados. |
| **Soporte en Go**           | Goroutines y canales.                                         | Distribución interna en múltiples hilos. |

---

## **📜 Conceptos Relacionados**

| **Término**       | **Definición**                                                                                             |
|--------------------|-----------------------------------------------------------------------------------------------------------|
| **Síncrono**       | Tareas ejecutadas una tras otra, bloqueando el flujo hasta que cada una termina.                          |
| **Asíncrono**      | Tareas que no bloquean el flujo; permiten avanzar mientras una tarea espera.                              |
| **Goroutine**      | Función concurrente ligera gestionada por el runtime de Go.                                               |
| **Canales**        | Herramientas para comunicar y sincronizar datos entre Goroutines de manera segura.                        |

---