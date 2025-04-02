package main

import (
	"fmt"
	"strconv"
)

// var era = "AD"

// type celcuis float64
// type kelvin float64

// func kelvinToCelcuis(k kelvin) celcuis {
// 	return celcuis(k - 273.15)
// }

// type kelvin float64

// func fakeSensor() kelvin {
// 	return kelvin(rand.Intn(151) + 150)
// }

// func realSensor() kelvin {
// 	return 0
// }

// func measureTemperature(samples int, sensor func() kelvin) {
// 	for i := 0; i < samples; i++ {
// 		k := sensor()

// 		fmt.Printf("%v K\n", k)
// 		time.Sleep(time.Second)
// 	}
// }

// var f = func() {
// 	fmt.Println("Dress up for the masquerade")
// }

// func terraform(planets [8]string) {
// 	for i := range planets {
// 		planets[i] = "New " + planets[i]
// 		fmt.Println(planets[i])
// 	}
// }

// func terraform(prefix string, worlds ...string) []string {
// 	nweWorlds := make([]string, len(worlds))

// 	for i := range worlds {
// 		nweWorlds[i] = prefix + " " + worlds[i]
// 	}
// 	return nweWorlds
// }

func main() {

	// var i int = 33
	// var s string = strconv.Itoa(i)

	var str string = "rohit"
	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)
	// fmt.Println(s)

	// var name string = "rohit"
	// // fmt.Printf("%T", name)
	// fmt.Printf("name is of type %T \n", name)

	// var name string
	// var num int

	// fmt.Print("Enter your name: ")
	// count, err := fmt.Scanf("%s %d", &name, &num)

	// if err != nil {
	// 	fmt.Println("Error inputing name", err)
	// }

	// fmt.Println(count)
	// fmt.Println("Your name is ", name, " and ", num)

	// type location struct {
	// 	Lat  float64 `json:"latitude"`
	// 	Long float64 `json:"longitude"`
	// }

	// curosity := location{435.45, 43543.4356345}

	// bytes, err := json.Marshal(curosity)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Println(string(bytes))

	// type location struct {
	// 	Lat, Long float64
	// }

	// curosity := location{.34, .343}
	// bytes, err := json.Marshal(curosity)

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Println(string(bytes))
	// var temp location
	// err = json.Unmarshal(bytes, &temp)

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Println(temp)

	// type location struct {
	// 	name string
	// 	lat  float64
	// 	long float64
	// }

	// locations := []location{
	// 	{name: "Bradbury Landing", lat: -4.5895, long: 137.4417},
	// 	{name: "Columbia Memorial Station", lat: -14.5684, long: 175.472636},
	// 	{name: "Challenger Memorial Station", lat: -1.9462, long: 354.4734},
	// }

	// type location struct {
	// 	lat, long float64
	// }

	// bradbury := location{-4.5, 234.44}
	// curosity := bradbury

	// curosity.lat = 30000
	// fmt.Println(curosity)
	// fmt.Println(bradbury)

	// opportunity := location{lat: -.34, long: 34.4}
	// fmt.Println(opportunity)

	// insight := location{lat: 4.5, long: 135.9}
	// fmt.Println(insight)

	// type location struct {
	// 	lat  float64
	// 	long float64
	// }

	// var spirit location
	// spirit.lat = 23
	// spirit.long = 34

	// fmt.Println(spirit)

	// var curosity struct {
	// 	lat  float64
	// 	long float64
	// }

	// curosity.lat = -.32423
	// curosity.long = -234.32423
	// fmt.Println(curosity)

	// var temperatures = []float64{
	// 	-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	// }

	// set := make(map[float64]bool)

	// for _, t := range temperatures {
	// 	set[t] = true
	// }

	// if set[-28.0] {
	// 	fmt.Println("is a member of set")
	// }

	// fmt.Println(set)

	// unique := make([]float64, 0, len(set))
	// for t := range set {
	// 	unique = append(unique, t)
	// }
	// sort.Float64s(unique)

	// fmt.Println(unique)

	// temperature := []float64{-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0}
	// groups := make(map[float64][]float64)

	// for _, t := range temperature {
	// 	g := math.Trunc(t/10) * 10
	// 	groups[g] = append(groups[g], t)
	// }

	// fmt.Println(groups)

	// temperature := []float64{-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0}
	// frequency := make(map[float64]int)

	// for _, t := range temperature {
	// 	frequency[t]++
	// }

	// fmt.Println(frequency)

	// temperature := map[string]int{
	// 	"Earth": 19,
	// 	"Mars":  -9,
	// }

	// newTemperature := temperature
	// newTemperature["Moon"] = -10
	// fmt.Println(temperature)
	// fmt.Println(newTemperature)

	// delete(newTemperature, "Moon")

	// fmt.Println(temperature)
	// fmt.Println(newTemperature)

	// temp := temperature["Earth"]

	// temperature["Earth"] = 100
	// temperature["Mars"] = 0
	// temperature["Moon"] = 10000

	// fmt.Println(temperature)
	// fmt.Println(temperature["Moon"])

	// if moon, ok := temperature["Moon"]; ok {
	// 	fmt.Println("Moon is present with temperature of", moon)
	// } else {
	// 	fmt.Println("Moon is not present")
	// }

	// planets := []string{"Venus", "Mars", "Jupiter"}
	// newPlanets := terraform("New", planets...)
	// fmt.Println(newPlanets)

	// dwarfs := make([]string, 0, 10)
	// dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	// fmt.Println(dwarfs)

	// dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	// dwarfs = append(dwarfs, "Orcus")
	// fmt.Println(dwarfs)

	// fmt.Println(len(dwarfs))
	// fmt.Println(cap(dwarfs))

	// dwarfs = append(dwarfs, "Rohit", "Pluto", "Haumea", "Makemake", "temping")
	// fmt.Println(len(dwarfs))
	// fmt.Println(cap(dwarfs))

	// dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	// // sort.StringSlice(dwarfs).Sort()
	// sort.Strings(dwarfs)
	// fmt.Println(dwarfs)

	// planets := [...]string{
	// 	"Mercury",
	// 	"Venus",
	// 	"Earth",
	// 	"Mars",
	// 	"Jupiter",
	// 	"Saturn",
	// 	"Uranus",
	// 	"Neptune",
	// }

	// terrestrial := planets[0:4]
	// gasGiants := planets[4:6]
	// iceGiants := planets[6:8]

	// fmt.Println(terrestrial)
	// fmt.Println(gasGiants)
	// fmt.Println(iceGiants)

	// terraform(planets)
	// fmt.Println(planets)

	// planets := [...]string{
	// 	"temp1", "temp2", "temp3", "temp4",
	// }

	// newPlanets := planets
	// planets[0] = "test1"

	// fmt.Println(planets)
	// fmt.Println(newPlanets)
	// fmt.Println(planets)

	// fmt.Println(len(planets))

	// for i := 0; i < len(planets); i++ {
	// 	fmt.Println(planets[i])
	// }

	// for i, planet := range planets {
	// 	fmt.Println(i, planet)
	// }

	// dwarfs := [5]string{"temp1", "temp2", "temp3", "temp4", "temp5"}
	// fmt.Println(dwarfs)

	// var planets [8]string
	// planets[0] = "temp1"
	// planets[1] = "temp2"
	// planets[2] = "temp3"

	// earth := planets[2]
	// fmt.Println(planets)
	// fmt.Println(earth)
	// fmt.Println(len(planets))

	// func(message string) {
	// 	fmt.Println("Hila ", message)
	// }("rohan")

	// f := func(message string) {
	// 	fmt.Println(message)
	// }

	// f("Hello rohan how u doin")

	// measureTemperature(3, fakeSensor)

	// sesor := fakeSensor
	// fmt.Println(sesor())

	// sesor = realSensor
	// fmt.Println(sesor())

	// var k kelvin = 354.5
	// c := kelvinToCelcuis(k)
	// fmt.Println(k, c)

	// year := 2018
	// switch month := rand.Intn(12) + 1; month {
	// case 2:
	// 	day := rand.Intn(28) + 1
	// 	fmt.Println(era, year, month, day)
	// case 4, 6, 9, 11:
	// 	day := rand.Intn(30) + 1
	// 	fmt.Println(era, year, month, day)
	// default:
	// 	day := rand.Intn(31) + 1
	// 	fmt.Println(era, year, month, day)
	// }

	// fmt.Print("My weight on the surface of mars is ")
	// fmt.Print(149.0 * 0.45)
	// fmt.Print(" lbs and i would be ")
	// fmt.Print(41 * 365.2425 / 687)
	// fmt.Print(" years old")

	// fmt.Printf("My weight on the surface of %v is %v lbs\n", "Earth", 170.45)

	// const lightSpeed = 299792
	// var distance = 5600000
	// fmt.Println(distance/lightSpeed, "seconds")
	// distance = 401000000
	// fmt.Println(distance/lightSpeed, "seconds")

	// var (
	// 	distance = 40000
	// 	speed    = 204045
	// )

	// const hoursPerDay, minutesPerHour = 24, 60

	// var num = rand.Intn(10) + 1
	// fmt.Println(num)

	// num = rand.Intn(10) + 1
	// fmt.Println(num)

	// fmt.Println("You find yourself in a dumly lit carven.")
	// var command = "walk outsidee"
	// var exit = strings.Contains(command, "r")
	// fmt.Println("You leave the cave:", exit)

	// var command = "go east"
	// if command == "go east" {
	// 	fmt.Println("Ypu head further up the mountain")
	// } else if command == "go inside" {
	// 	fmt.Println("You enter the cave where you live out the rest of your life")
	// } else {
	// 	fmt.Println("Didn't quitee get thtat")
	// }

	// var year = 2100
	// var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	// if leap {
	// 	fmt.Println("Look before you leap")
	// } else {
	// 	fmt.Println("Keep your feeet on the ground")
	// }

	// var command = "go inside"
	// switch command {
	// case "go east":
	// 	fmt.Println("You head further up the mountain.")
	// case "enter cave", "go inside":
	// 	fmt.Println("You find yourself in a dimly lit cavern.")
	// case "read sign":
	// 	fmt.Println("The sign reads 'No Minors'.")
	// default:
	// 	fmt.Println("Didn't quite get that.")
	// }

	// var room = "lake"
	// switch {
	// case room == "cave":
	// 	fmt.Println("You find yourself ina diimly lit carcen")
	// case room == "lake":
	// 	fmt.Println("the ice seems solid enough")
	// 	fallthrough
	// case room == "underwater":
	// 	fmt.Println("The water is freexing cold")

	// }

	// var count = 10
	// for count > 0 {
	// 	fmt.Println(count)
	// 	time.Sleep(time.Second)
	// 	count--
	// }
	// fmt.Println("Liftoff")

	// const num = 40
	// for {
	// 	var guessedNum = rand.Intn(100) + 1
	// 	fmt.Printf("Guessed number %v \n", guessedNum)
	// 	if guessedNum == num {
	// 		fmt.Println("You guessed the right number")
	// 		break
	// 	} else if guessedNum > num {
	// 		fmt.Println("guessed number is greater than actual number")
	// 	} else {
	// 		fmt.Println("guessed number is lesser than actual numbeer")
	// 	}
	// }

	// var count = 100
	// tem := 33

	// var count = 0
	// for count = 10; count > 0; count-- {
	// 	fmt.Println(count)
	// }
	// fmt.Println(count)
	// for count := 10; count > 0; count-- {
	// 	fmt.Println(count)
	// }

	// if num := rand.Intn(3); num == 0 {
	// 	fmt.Println("Space adventures")
	// } else if num == 1 {
	// 	fmt.Println("spacex")
	// } else {
	// 	fmt.Println("Virgin galctic")
	// }

	// switch num := rand.Intn(10); num {
	// case 0:
	// 	fmt.Println("space advcentures")
	// case 1:
	// 	fmt.Println("spacex")
	// default:
	// 	fmt.Println("nothing")
	// }

	// a := "text"
	// fmt.Printf("Typee %T for %[1]v\n", a)

	// b := 34
	// fmt.Printf("Typee %T for %[1]v\n", b)

	// c := 4.12
	// fmt.Printf("Typee %T for %[1]v\n", c)

	// d := true
	// fmt.Printf("Typee %T for %[1]v\n", d)

	// var distance int64 = 42.3e12
	// fmt.Println("Alpha centauri is ", distance, " km away")

	// fmt.Println(`hello
	// whow u doning
	// rohit`)

	// message := "shalom"
	// c := string(message[5])
	// fmt.Println(c)
	// fmt.Println(len(message))
	// upperchar := strings.ToUpper(c)
	// fmt.Println(upperchar)

	// question := "helo mann"
	// for _, c := range question {
	// 	// fmt.Println(string(c))
	// 	fmt.Printf("%c", c)
	// }

	// kelvin := 235.6
	// celcius := kelvinToCelcuis(kelvin)
	// fmt.Print(kelvin, "deghree K is ", celcius, "degree  c")

}

// func kelvinToCelcuis(k float64) float64 {
// 	k -= 273.15
// 	return k
// }
