# Weather Prediction
![go:{version}-min](https://img.shields.io/badge/go-1.23--mini-blue.svg)
## Problema
En una galaxia lejana, existen tres civilizaciones: Vulcanos, Ferengis y Betazoides. Cada civilización vive en paz en su respectivo planeta.
Dominan la predicción del clima mediante un complejo sistema informático.
A continuación el diagrama del sistema solar

### Premisas
- El planeta Ferengi se desplaza con una velocidad angular de 1 grados/día en sentido horario. Su distancia con respecto al sol es de 500Km.
- El planeta Vulcano se desplaza con una velocidad angular de 5 grados/día en sentido antihorario. Su distancia con respecto al sol es de 1000Km.
- El planeta Betazoide se desplaza con una velocidad angular de 3 grados/día en sentido horario. Su distancia con respecto al sol es de 2000Km.
- Todas las órbitas son circulares.

Cuando los tres planetas están alineados entre sí y a su vez alineados con respecto al sol, el sistema solar experimenta un período de sequía.

Cuando los tres planetas no están alineados, forman entre sí un triángulo. Es sabido que en el momento en el que el sol se encuentra dentro del triángulo, el sistema solar experimenta un período de lluvia, teniendo este un pico de intensidad cuando el perímetro del triángulo está en su máximo. Si el sol NO se encuentra dentro del triángulo, el sistema solar experimenta un período de normalidad.

Las condiciones óptimas de presión y temperatura se dan cuando los tres planetas están alineados entre sí pero no están alineados con el sol.

### Entregable
- Se debe desarrollar un programa que pueda predecir la siguiente información en los próximos X años:

    - ¿Cuántos períodos de sequía habrá?
    - ¿Cuántos períodos de lluvia habrá y qué día será el pico máximo de lluvia
    - ¿Cuántos períodos de condiciones óptimas de presión y temperatura habrá?
- Para poder utilizar el sistema como un servicio a las otras civilizaciones, los Vulcanos requieren tener una base de datos con las condiciones meteorológicas de todos los días y brindar una API REST de consulta sobre las condiciones de un día en particular.

## Solución
Cada planeta tiene un movimiento circular uniforme definido por:

- Su velocidad angular (°/día)

- Su distancia al sol (radio de la órbita)

- El sentido de rotación (horario o antihorario)

La posición (x, y) de un planeta, en un día dado, se calcula usando:

- La distancia del planeta al sol

- El ángulo recorrido en radianes

El ángulo se convierte de grados a radianes porque las funciones trigonométricas en Go trabajan en radianes.

### Detección de alineaciones

Se verifica si tres puntos están alineados usando el determinante de las coordenadas, equivalente al área del triángulo formado por los puntos.

Si el valor absoluto del determinante es cercano a cero (con un umbral de tolerancia), los puntos están alineados.

### Contención del sol dentro del triángulo

Para determinar si el sol (0, 0) está dentro del triángulo formado por los tres planetas, se calcula el área total del triángulo y se compara con la suma de las áreas de los subtriángulos que incluyen al sol.

Si la suma de las subáreas es igual al área total, el sol está dentro del triángulo.

### Perimetro del triangulo

El perímetro del triángulo formado por los tres planetas se calcula sumando las distancias entre sus vértices.

## Implementación
- Recorre cada dia del numero de años que se solicite
- Calcula las posiciones de los tres planetas
- Verifica si los planetas están alineados
    - Si están alineados entre sí y con el sol, se registra un período de sequía.
    - Si están alineados entre sí pero no con el sol, se registra un período de condiciones óptimas.
- Si no están alineados, verifica si el sol está dentro del triángulo formado por los planeta.
    - Si el sol está dentro, se registra un período de lluvia.
    - Calcula el perímetro del triángulo para determinar el día de máxima lluvia.

## Requisitos

- Go 1.23
- MySQL (Opcional)

## Configuración local

- Clonación del repositorio
```bash
git clone git@github.com:Juram09/Weather-Predictor.git
```
- Configuración de base de datos local (Opcional)

  Esto para la consulta y guardado de las condiciones meteorologicas por dias.

    - En caso de requerirlo, abrir el terminal de MySQL
    ```bash
    mysql -u root -p
    ```

    - Creación de base de datos

    ```bash
    CREATE DATABASE weather;
    USE weather;
    ```
    - Creación de la tabla de clima

    ```bash
    CREATE TABLE weather (id INT PRIMARY KEY, meteorolic_conditions VARCHAR(255));
    ```
- Ejecución de programa
```bash
go run main.go
```

## Uso

### Obtener numero de periodos de sequía
- **GET** localhost:8080/weather/drought
> Todas las peticiones tienen que tener el query parameter "years" seteado como un entero positivo.
```bash
curl --location 'localhost:8080/weather/drought?years=10'
```

#### Respuestas

***200***

```json
{
    "years": 10,
    "droughts": 40
}
```

***400***

```json
{
    "error": "Invalid years"
}
```

### Obtener numero de periodos de lluvía
- **GET** localhost:8080/weather/rainy
> Todas las peticiones tienen que tener el query parameter "years" seteado como un entero positivo.
```bash
curl --location 'localhost:8080/weather/rainy?years=10'
```

#### Respuestas

***200***

```json
{
    "years": 10,
    "rains": 1208,
    "rain_peak": 72
}
```

***400***

```json
{
    "error": "Invalid years"
}
```

### Obtener numero de periodos de condiciones optimas
- **GET** localhost:8080/weather/optimal
> Todas las peticiones tienen que tener el query parameter "years" seteado como un entero positivo.
```bash
curl --location 'localhost:8080/weather/optimal?years=10'
```

#### Respuestas

***200***

```json
{
    "years": 10,
    "optimal": 41
}
```

***400***

```json
{
    "error": "Invalid years"
}
```

### Obtener condiciones meteorologicas de un dia en concreto
- **GET** localhost:8080/weather
> Todas las peticiones tienen que tener el query parameter "day" seteado como un entero positivo.
```bash
curl --location 'localhost:8080/weather?day=1234'
```

#### Respuestas

***200***

```json
{
    "day": 1234,
    "weather": "rain"
}
```

***400***

```json
{
    "error": "Invalid day"
}
```
