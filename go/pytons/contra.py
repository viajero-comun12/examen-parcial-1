import sys
import time
import itertools
import matplotlib.pyplot as plt
from main import authenticate_user

letras = ["a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"]
letrasmay = ["A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"]
numeros = ["1","2","3","4","5","6","7","8","9"]
caracteres = [".","; ",",","!","$","#","%","&","/","(",")","=","?","¿"]
almacen = []
inte = []

alfabeto = letras + letrasmay + numeros + caracteres

intentos = 0
encontrada = None

tiempo_inicio = None

def rompecontraseñas(username, password=None, max_len=4):
    global intentos, encontrada, tiempo_inicio
    intentos = 0
    encontrada = None
    tiempo_inicio = time.perf_counter()
    almacen.clear()
    inte.clear()

    if isinstance(password, str):
        for ch in password:
            if ch not in alfabeto:
                print("El carácter", ch, " no está en el alfabeto definido. Agrega ese carácter al alfabeto o cambia la contraseña.")
        target_mode = True
    else:
        target_mode = False

    A = len(alfabeto)
    total_teorico = sum(A ** i for i in range(1, max_len + 1))
    print(f"Alfabeto size={A}, max_len={max_len}, combinaciones teóricas={total_teorico:,}")
    for longitud in range(1, max_len + 1):
        for combo in itertools.product(alfabeto, repeat=longitud):
            intentos += 1
            intento = "".join(combo)

     
            if target_mode:
                candidato = intento
                ok_local = (candidato == password)
            else:

                candidato = intento
                try:
                    ok_local = authenticate_user(username, candidato)
                except Exception as e:
                    print("Error en authenticate_user:", e)

                    tiempo_fin = time.perf_counter()
                    tiempo_transcurrido = tiempo_fin - tiempo_inicio
                    return {"tiempo": tiempo_transcurrido}

            if ok_local:
                encontrada = candidato
                tiempo_fin = time.perf_counter()
                tiempo_transcurrido = tiempo_fin - tiempo_inicio
                print("Contraseña encontrada")
                print("Contraseña:", encontrada)
                print("Intentos:", intentos)
                inte.append(intentos)
                almacen.append(tiempo_transcurrido)

                return {tiempo_transcurrido}

            if intentos % 100000 == 0:
                tiempo_parcial = time.perf_counter() - tiempo_inicio
                tasa = intentos / tiempo_parcial if tiempo_parcial > 0 else 0
                print(f"Progreso: intentos={intentos:,}  tiempo={tiempo_parcial:.1f}s  tasa={tasa:.0f} intentos/s")

    tiempo_fin = time.perf_counter()
    tiempo_transcurrido = tiempo_fin - tiempo_inicio
    print("No se encontró la contraseña.")
    print("Intentos:", intentos)
    print("Tiempo transcurrido:", tiempo_transcurrido)
    almacen.append(tiempo_transcurrido)
    inte.append(intentos)
    return {"found": None, "attempts": intentos, "time": tiempo_transcurrido}


def mostrar_grafico():
    if not almacen or not inte:
        print("No hay datos para graficar.")
        return
    plt.figure(figsize=(6,4))
    plt.plot(almacen, inte, marker='o')
    plt.xscale("linear")
    plt.yscale("log")
    plt.xlabel("Tiempo (s)")
    plt.title("Simulación fuerza bruta (tiempo vs intentos)")
    plt.grid(True)
    plt.show()
def main():
    if len(sys.argv) < 3:
        print("Uso: python local_bruteforce.py <username> <max_len>")
        print("Ejemplo: python local_bruteforce.py admin 4")
        sys.exit(1)

    username = sys.argv[1]
    max_len = int(sys.argv[2])

    print("Iniciando simulación para usuario"+ username)

    res = rompecontraseñas(username, password=None, max_len=max_len)
    print("timepo", res)

    mostrar_grafico()

if __name__ == "__main__":
    main()