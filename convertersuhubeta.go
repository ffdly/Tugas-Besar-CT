package main


import "fmt"

func main() {
	var pilihan int
	var suhu float64

	fmt.Println("=== Converter Suhu ===")
	fmt.Println("1. Celcius ke Fahrenheit")
	fmt.Println("2. Celcius ke Kelvin")
	fmt.Println("3. Fahrenheit ke Celcius")
	fmt.Println("4. Kelvin ke Celcius")
	fmt.Print("Pilih menu (1-4): ")
	fmt.Scan(&pilihan)

	if pilihan == 0 {
		fmt.println("program kalkulasi selesai.")
		break
	}

	fmt.Print("Masukkan nilai suhu: ")
	fmt.Scan(&suhu)

	switch pilihan {
	case 1:
		fmt.Printf("Hasil: %.2f °F\n", celciusToFahrenheit(suhu))
	case 2:
		fmt.Printf("Hasil: %.2f K\n", celciusToKelvin(suhu))
	case 3:
		fmt.Printf("Hasil: %.2f °C\n", fahrenheitToCelcius(suhu))
	case 4:
		fmt.Printf("Hasil: %.2f °C\n", kelvinToCelcius(suhu))
	default:
		fmt.Println("Pilihan tidak valid")
	}
}

// Fungsi konversi
func celciusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

func celciusToKelvin(c float64) float64 {
	return c + 273.15
}

func fahrenheitToCelcius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func kelvinToCelcius(k float64) float64 {
	return k - 273.15
}

