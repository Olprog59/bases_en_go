package main

import (
	"fmt"
	"strings"
)

type Post struct {
	Titre       string
	Description string
	published   bool
}

func (p Post) HeadLine() string {
	return fmt.Sprintf("%v - %v\n", p.Titre, p.Description[:50])
}

func (p Post) Published() bool {
	return p.published
}

func (p *Post) publish() {
	p.published = true
}

func (p *Post) unPublish() {
	p.published = false
}

func main() {
	// post := &Post{
	post := Post{
		Titre: "Félicitations pour un nouveau poste attendu",
		Description: `Bonjour Jean, 
Je suis très content que tu aies enfin le poste que tu esperais depuis quelques année 
maintenant. Je sais que tu es fin prêt pour assumer un tel poste et que tu pourras 
t'épanouir pleinnement dans cette nouvelle société. Je te félicite pour ce nouveau
poste et suis très impressionné par le parcours que tu as suvi jusqu'à présent. 
Je sais que tu n'as pas fini de grimper et te souhaite le meilleur dans cette nouvelle
étape de ta vie professionnelle. Bonne chance et bon courage pour ta prise de poste. 
N'hésite pas à me donner des nouvelles dans quelques semaines. 
Félicitations ! 
Francis`,
		published: false,
	}

	fmt.Printf("Mon post est : %v\n", post.HeadLine())

	post.publish()

	//fmt.Printf("%#+v\n", post)

	//UpperCase(post)
	UpperCase(&post)
	fmt.Println(post.Titre)
}

func UpperCase(post *Post) {
	post.Titre = strings.ToUpper(post.Titre)
}
