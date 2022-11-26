package main


import "fmt"

type SerVivo interface {
	estaVivo() bool
}
type humano interface{
	respirar()
	pensar()
	comer()
	sexo() string
}
type animal interface{
	respirar()
	comer()
	EsCarnivoro() bool
	
}
type vegetal interface{
	ClasificacionVegetal() string
}
// genero humano

type hombre struct{
	edad int
	altura float32
	peso float32
	respirando bool
	pensando bool
	comiendo bool
	esHombre bool
	vivo bool
}

type mujer struct{
	hombre
}
func (h *hombre) respirar() {
	h.respirando = true
}
func (h *hombre) pensar() {
	h.pensando = true
}
func (h *hombre) comer() {
	h.comiendo = true
}
func (h *hombre) sexo() string {
	if h.esHombre{
		return "Hombre"
	} else {
		return "Mujer"
	}
}
func (h *hombre) estaVivo() bool {
	return h.vivo
}

func HumanosRespirando(hu humano){
	hu.respirar()
	fmt.Printf("soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

// genero animal

type perro struct{
	respirando bool
	comiendo bool
	Carnivoro bool
	vivo bool
}
func (p *perro) respirar() {
	p.respirando = true
}
func (p *perro) comer() {
	p.comiendo = true
}
func (p *perro) EsCarnivoro() bool {
	return p.Carnivoro
}
func (p *perro) estaVivo() bool { return p.vivo}
func AnimalesRespirando(an animal){
	an.respirar()
	fmt.Printf("soy un/a perro, y ya estoy respirando \n")
}
func AnimalesCarnivoros(an animal){
	carnivoroEs:=an.EsCarnivoro()
	if carnivoroEs==true{
		fmt.Println("Soy carnivoro")
	} else {
		fmt.Println("No soy carnivoro")
	}
}

/* SerVivo */

func estoyVivo(v SerVivo) bool {
	return v.estaVivo()
}
func main(){
	pedro:=new(hombre)
	pedro.esHombre = true
	HumanosRespirando(pedro)
	pedro.vivo=true
	fmt.Printf("Estoy vivo? %t \n", estoyVivo(pedro))

	ana:=new(mujer)
	HumanosRespirando(ana)
	ana.vivo=true
	fmt.Printf("Estoy vivo? %t \n", estoyVivo(ana))

	firulais:=new(perro)
	firulais.Carnivoro = true
	firulais.vivo=true
	AnimalesRespirando(firulais)
	AnimalesCarnivoros(firulais)

	fmt.Printf("Estoy vivo? %t \n", estoyVivo(firulais))
}
