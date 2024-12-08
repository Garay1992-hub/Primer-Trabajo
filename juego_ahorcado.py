import random

class JuegoAhorcado:
    def __init__(self):
        # Banco de palabras
        self.palabras = [
            "python", "programacion", "computadora", 
            "desarrollo", "inteligencia", "tecnologia"
        ]
        # Selección aleatoria de palabra
        self.palabra_secreta = random.choice(self.palabras)
        
        # Estado del juego
        self.letras_adivinadas = set()
        self.intentos_maximos = 6
        self.intentos_restantes = self.intentos_maximos

    def mostrar_palabra_actual(self):
        """Muestra el estado actual de la palabra"""
        return ' '.join(
            letra if letra in self.letras_adivinadas else '_' 
            for letra in self.palabra_secreta
        )

    def intentar_letra(self, letra):
        """Procesa el intento de adivinar una letra"""
        letra = letra.lower()
        
        # Validaciones
        if letra in self.letras_adivinadas:
            return "Letra ya intentada"
        
        self.letras_adivinadas.add(letra)
        
        # Verificar si la letra está en la palabra
        if letra not in self.palabra_secreta:
            self.intentos_restantes -= 1
            return "Letra incorrecta"
        
        return "Letra correcta"

    def verificar_victoria(self):
        """Comprueba si el jugador ha adivinado la palabra"""
        return all(letra in self.letras_adivinadas for letra in self.palabra_secreta)

    def jugar(self):
        """Ciclo principal del juego"""
        print("¡Bienvenido al Juego del Ahorcado!")
        
        while self.intentos_restantes > 0:
            print(f"\nIntentos restantes: {self.intentos_restantes}")
            print("Palabra:", self.mostrar_palabra_actual())
            
            # Entrada de letra
            letra = input("Ingresa una letra: ").lower()
            
            # Procesar intento
            resultado = self.intentar_letra(letra)
            print(resultado)
            
            # Verificar victoria
            if self.verificar_victoria():
                print(f"\n¡Ganaste! La palabra era: {self.palabra_secreta}")
                return
            
            # Verificar derrota
            if self.intentos_restantes == 0:
                print(f"\n¡Perdiste! La palabra era: {self.palabra_secreta}")

def main():
    """Función principal para iniciar el juego"""
    juego = JuegoAhorcado()
    juego.jugar()

# Punto de entrada del programa
if __name__ == "__main__":
    main()