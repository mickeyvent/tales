package main

import(
	"fmt"
	"time"
	"os"
	"os/exec"
	"github.com/liamg/tml"
	"github.com/briandowns/spinner"
)
var(
	dt,ps string
	a,b,c,z,anony float64
)
func main(){
	for po:=0; po>=0; po++{
		clear()
		table()
		dt = " "
		tml.Printf("<bg-magenta><black>Tales<bg-black><red><blink>:")
		fmt.Scanf("%v",&dt)
		if dt == "a" || dt == "A"{
			fmt.Printf("Valor de A:")
			fmt.Scanf("%v",&a)
		}else if dt == "b" || dt == "B"{
			fmt.Printf("Valor de B:")
			fmt.Scanf("%v",&b)
		}else if dt == "c" || dt == "C"{
			fmt.Printf("Valor de C:")
			fmt.Scanf("%v",&c)
		}else if dt == "run" || dt == "next"{
			segundo()
		}else if dt == "help"{
			help()
		}else if dt == "exit"{
			os.Exit(0)
		}else if dt == "clear"{
			a = 0
			b = 0
			c = 0
			z = 0
			anony = 0
		}else{
			fmt.Printf("No se reconoce: %v",dt)
			time.Sleep(2*time.Second)
		}
	}
}
func segundo(){
	for bu := 0; bu >=0; bu++{
		if a > 0{
			if b > 0{
				if c > 0{
					z = a
					a = c * a
					anony = a / b
					a = z
					break
				}else{
					tml.Printf("<red>No hay ningun dato del lado <green>C \n")
					time.Sleep(2*time.Second)
					break
				}
			}else{
				tml.Printf("<red>No hay ningun dato del lado <green>B \n")
				time.Sleep(2*time.Second)
				break
			}
		}else{
			tml.Printf("<red>No hay ninguna dato del lado <green>A \n")
			time.Sleep(2*time.Second)
			break
		}
	}
}
func help(){
	clear()
	animate()
	fmt.Printf("Tales es usa el teorema de tales para obtener la proporcion.\n")
	fmt.Printf("para modificar una variable solo escribe la letra y \n")
	fmt.Printf("despues solo ingresa el valor de esa variable. ejemplo:\n")
	fmt.Printf("  $Tales: A\n")
	fmt.Printf("     ingresa valor de A: 23\n")
	fmt.Printf("despues para de ingresar todos los datos de A,B,C\n")
	fmt.Printf("solo escriba Run y obtendra el resultado.\n")
	pause()
}
func animate(){
	s := spinner.New(spinner.CharSets[21],100 * time.Millisecond)
	s.Suffix = " Help !!! Help!!!"
	s.Start()
	time.Sleep(2*time.Second)
	s.Stop()
}
func pause(){
	fmt.Printf("\nENTER PARA CONTINUAR:")
	fmt.Scanf("%v",&ps)
}
func table(){
	fmt.Printf("     TEOREMA DE TALES... \n")
	tml.Printf("           <green>A<white>(%v)     <green>?<white>(%v)\n",a,anony)
	fmt.Printf("          _______ = ________\n")
	tml.Printf("           <green>B<white>(%v)     <green>C<white>(%v)\n",b,c)
	tml.Printf("     [<yellow>help<white>] [<yellow>clear<white>] [<yellow>run<white>] [<yellow>exit<white>]\n")
}
func clear(){
	cl := exec.Command("clear")
	cl.Stdout = os.Stdout
	cl.Run()
}
