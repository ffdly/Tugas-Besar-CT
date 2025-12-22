package main

import "fmt"

func main() {
	var pilihan int
	var suhu float64

	for {
		fmt.Println("\n=== CONVERTER SUHU ===")
		fmt.Println("=== DARI CELCIUS ===")
		fmt.Println("1.  Celcius ke Fahrenheit")
		fmt.Println("2.  Celcius ke Kelvin")
		fmt.Println("3.  Celcius ke Reamur")

		fmt.Println("\n=== DARI FAHRENHEIT ===")
		fmt.Println("4.  Fahrenheit ke Celcius")
		fmt.Println("5.  Fahrenheit ke Kelvin")
		fmt.Println("6.  Fahrenheit ke Reamur")

		fmt.Println("\n=== DARI KELVIN ===")
		fmt.Println("7.  Kelvin ke Celcius")
		fmt.Println("8.  Kelvin ke Fahrenheit")
		fmt.Println("9.  Kelvin ke Reamur")

		fmt.Println("\n=== DARI REAMUR ===")
		fmt.Println("10. Reamur ke Celcius")
		fmt.Println("11. Reamur ke Fahrenheit")
		fmt.Println("12. Reamur ke Kelvin")

		fmt.Println("\n0.  Keluar")
		fmt.Print("Pilih menu (0-12): ")

		
		for {
			fmt.Scan(&pilihan)
			
			if pilihan == 0 {
				fmt.Println("Kalkulasi Selesai.")
				return 
			}
			
			if pilihan >= 1 && pilihan <= 12 {
				break 
			}
			fmt.Println("Pilihan tidak valid! Masukkan angka 0-12.")
			fmt.Print("Pilih menu (0-12): ")
		}

		fmt.Print("Masukkan nilai suhu: ")
		fmt.Scan(&suhu)

		switch pilihan {
		
		case 1:
			fmt.Printf("Hasil: %.2f °F\n", celciusToFahrenheit(suhu))
		case 2:
			fmt.Printf("Hasil: %.2f K\n", celciusToKelvin(suhu))
		case 3:
			fmt.Printf("Hasil: %.2f °Ré\n", celciusToReamur(suhu))

		
		case 4:
			fmt.Printf("Hasil: %.2f °C\n", fahrenheitToCelcius(suhu))
		case 5:
			fmt.Printf("Hasil: %.2f K\n", fahrenheitToKelvin(suhu))
		case 6:
			fmt.Printf("Hasil: %.2f °Ré\n", fahrenheitToReamur(suhu))

		
		case 7:
			fmt.Printf("Hasil: %.2f °C\n", kelvinToCelcius(suhu))
		case 8:
			fmt.Printf("Hasil: %.2f °F\n", kelvinToFahrenheit(suhu))
		case 9:
			fmt.Printf("Hasil: %.2f °Ré\n", kelvinToReamur(suhu))

		
		case 10:
			fmt.Printf("Hasil: %.2f °C\n", reamurToCelcius(suhu))
		case 11:
			fmt.Printf("Hasil: %.2f °F\n", reamurToFahrenheit(suhu))
		case 12:
			fmt.Printf("Hasil: %.2f K\n", reamurToKelvin(suhu))
		}
	}
}


func celciusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

func celciusToKelvin(c float64) float64 {
	return c + 273.15
}

func celciusToReamur(c float64) float64 {
	return c * 4 / 5
}


func fahrenheitToCelcius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func fahrenheitToKelvin(f float64) float64 {
	return celciusToKelvin(fahrenheitToCelcius(f))
}

func fahrenheitToReamur(f float64) float64 {
	return celciusToReamur(fahrenheitToCelcius(f))
}


func kelvinToCelcius(k float64) float64 {
	return k - 273.15
}

func kelvinToFahrenheit(k float64) float64 {
	return celciusToFahrenheit(kelvinToCelcius(k))
}

func kelvinToReamur(k float64) float64 {
	return celciusToReamur(kelvinToCelcius(k))
}


func reamurToCelcius(r float64) float64 {
	return r * 5 / 4
}

func reamurToFahrenheit(r float64) float64 {
	return celciusToFahrenheit(reamurToCelcius(r))
}

func reamurToKelvin(r float64) float64 {
	return celciusToKelvin(reamurToCelcius(r))
}
