# Go Server

Mantiene una cola de trabajadores para poder ejecutar procesos pesados que tardan mucho tiempo, de manera concurrente.

Por ahora solo ejecuta codigos Javascript usando Node (ver ejemplo).

## Compilar

```bash
go build -o server *.go
```

## Ejecutar servidor

El argumento `-n` es la cantidad de trabajadores maximo.

```
./server -n 3
```

## Agregar peticiones a la cola

Este ejemplo manda un gran numero de peticiones al servidor, los cuales son encoladas y ejecutadas cuando los trabajadores terminan su trabajo anterior.

```
```

Con json

```json
{
	"code": "var a=5; for(i=0;i<500000000;i++) a++; console.log(a);"
}
```
